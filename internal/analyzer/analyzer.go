package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"strconv"
	"strings"

	"github.com/shtemisu/loglinter/internal/rules"
	"github.com/shtemisu/loglinter/internal/stack"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

var percentFormatRe = regexp.MustCompile(`%%|%[+\-# 0]*\d*(?:\.\d*)?[a-zA-Z]`)

func stripFormatVerbs(s string) string {
	return percentFormatRe.ReplaceAllStringFunc(s, func(m string) string {
		if m == "%%" {
			return "%"
		}
		return ""
	})
}

type validator struct {
	Pass  *analysis.Pass
	Stack *stack.Stack[string]
}

func unquoting(l *ast.BasicLit) string {
	if l.Kind == token.STRING {
		s, err := strconv.Unquote(l.Value)
		if err != nil {
			return ""
		}
		return s
	}
	return ""
}

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "loglinter",
		Doc:      "Empty",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func isFormatFunction(name string) bool {
	return strings.HasSuffix(name, "f")
}

func getMethodName(c *ast.CallExpr) string {
	sel, ok := c.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	return sel.Sel.Name
}

func getPackageName(c *ast.CallExpr) string {
	sel, ok := c.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	if x, ok := sel.X.(*ast.Ident); ok {
		return x.Name
	}
	return ""
}

func (v *validator) checkLogMessage(n string, c *ast.CallExpr, val int) {
	switch n {
	case "log", "slog":
		v.checkCallLogSnap(c, val)
	case "zap":
		v.checkCallZap(c, val)
	default:
	}
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	var lastFile *token.File
	v := validator{
		Pass:  pass,
		Stack: stack.NewStack[string](),
	}
	nodeTypes := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Nodes(nodeTypes, func(n ast.Node, push bool) bool {
		currentFile := pass.Fset.File(n.Pos())

		if lastFile == nil || currentFile != lastFile {
			v.Stack.Clear()
			lastFile = currentFile
		}

		fnObject := typeutil.Callee(pass.TypesInfo, n.(*ast.CallExpr))
		funcTarget, ok := fnObject.(*types.Func)
		if !ok {
			return true
		}
		val, ok := messageArgIndex[funcTarget.FullName()]
		if ok {
			if push {
				v.checkCallLogSnap(n.(*ast.CallExpr), val)
			}
		}

		return true
	})
	return nil, nil
}

func (v *validator) checkCallLogSnap(call *ast.CallExpr, index int) {
	if len(call.Args) <= index {
		return
	}
	lit, ok := call.Args[index].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}
	msg := unquoting(lit)
	methodName := getMethodName(call)
	format := isFormatFunction(methodName)
	if format {
		msg = stripFormatVerbs(msg)
	}
	applyRules(msg, call, *v)
}

func (v *validator) checkCallZap(call *ast.CallExpr, index int) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	methodName := sel.Sel.Name

	logMethods := map[string]bool{
		"Debug": true, "Info": true, "Warn": true, "Error": true,
		"DPanic": true, "Panic": true, "Fatal": true,
		"Debugf": true, "Infof": true, "Warnf": true, "Errorf": true,
		"DPanicf": true, "Panicf": true, "Fatalf": true,
	}
	if !logMethods[methodName] {
		return
	}

	if len(call.Args) == 0 {
		return
	}
	logArg := call.Args[index]
	var msg string
	switch arg := logArg.(type) {
	case *ast.CallExpr:
		msg = checkCallError(arg)

	case *ast.Ident:
		msg = v.checkValueIdent(arg)

	case *ast.BasicLit:
		if arg.Kind == token.STRING {
			msg = unquoting(arg)
		}
	}

	if msg != "" {
		applyRules(msg, call, *v)
	}
}

func checkCallError(expr ast.Expr) string {

	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return ""
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	pkg, ok := sel.X.(*ast.Ident)
	if !ok || pkg.Name != "errors" {
		return ""
	}
	if sel.Sel.Name != "New" {
		return ""
	}

	if len(call.Args) == 0 {
		return ""
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return ""
	}
	return unquoting(lit)
}

func (v *validator) checkValueIdent(ident *ast.Ident) string {
	obj := v.Pass.TypesInfo.ObjectOf(ident)
	if obj == nil {
		return ""
	}

	for _, file := range v.Pass.Files {
		var found bool
		var result string
		ast.Inspect(file, func(n ast.Node) bool {
			if found {
				return false
			}
			switch node := n.(type) {
			case *ast.AssignStmt:
				for i, lhs := range node.Lhs {
					lhsIdent, ok := lhs.(*ast.Ident)
					if !ok {
						continue
					}
					if v.Pass.TypesInfo.ObjectOf(lhsIdent) != obj {
						continue
					}
					if i < len(node.Rhs) {
						if msg := checkCallError(node.Rhs[i]); msg != "" {
							result = msg
							found = true
							return false
						}
					}
				}

			case *ast.ValueSpec:
				for i, name := range node.Names {
					if v.Pass.TypesInfo.ObjectOf(name) != obj {
						continue
					}

					if i < len(node.Values) {
						if msg := checkCallError(node.Values[i]); msg != "" {
							result = msg
							found = true
							return false
						}
					}
				}
			}

			return true
		})
		if found {
			return result
		}
	}

	return ""
}

func applyRules(msg string, call *ast.CallExpr, v validator) {
	if !rules.IsLower(msg) {
		v.Pass.Reportf(call.Pos(), "log-message must be start lowercase char")
	}
	if !rules.OnlyEnglish(msg) {
		v.Pass.Reportf(call.Pos(), "log-message must be in English language")
	}
	if rules.HasSpecialChars(msg) {
		v.Pass.Reportf(call.Pos(), "log-message must not contains spec chars")
	}
	if rules.HasSensitiveData(msg) {
		v.Pass.Reportf(call.Pos(), "log-message must not contains sensetive data")
	}
}
