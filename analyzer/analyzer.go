package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/shtemisu/loglinter/internal/rules"
	"github.com/shtemisu/loglinter/internal/stack"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type validator struct {
	Pass  *analysis.Pass
	Stack *stack.Stack[string]
}

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "loglinter",
		Doc:      "Empty",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
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
		(*ast.FuncDecl)(nil), // добавляем, чтобы отслеживать функции
		(*ast.FuncLit)(nil),  // и анонимные функции
	}

	inspect.Nodes(nodeTypes, func(n ast.Node, push bool) bool {

		currentFile := pass.Fset.File(n.Pos())

		if lastFile == nil || currentFile != lastFile {
			v.Stack.Clear()
			lastFile = currentFile
		}

		if push {
			switch node := n.(type) {
			case *ast.FuncDecl:
				v.Stack.Push(node.Name.Name)
			case *ast.FuncLit:
				v.Stack.Push("anonymous")
			case *ast.CallExpr:
				v.checkCallLogSnap(node)
				v.checkCallZap(node)
				//	ast.Print(pass.Fset, n)
			}
		} else {
			switch n.(type) {
			case *ast.FuncDecl, *ast.FuncLit:
				v.Stack.Pop()
			}
		}
		return true
	})
	return nil, nil
}

func (v *validator) checkCallLogSnap(call *ast.CallExpr) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	if x, ok := sel.X.(*ast.Ident); ok && (x.Name == "log" || x.Name == "slog") {
		if len(call.Args) == 0 {
			return
		}
		lit, ok := call.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}
		msg := strings.Trim(lit.Value, "\"`")
		applyRules(msg, call, *v)
	}
}

func (v *validator) checkCallZap(call *ast.CallExpr) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	x, ok := sel.X.(*ast.Ident)
	if !ok || x.Name != "zap" {
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
	firstArg := call.Args[0]
	var msg string
	switch arg := firstArg.(type) {
	case *ast.CallExpr:
		msg = checkCallError(arg)

	case *ast.Ident:
		msg = v.checkValueIdent(arg)

	case *ast.BasicLit:
		if arg.Kind == token.STRING {
			msg = strings.Trim(arg.Value, "\"`")
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

	return strings.Trim(lit.Value, "\"`")
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
	if !rules.OnlyEnglishAndWithoutSpecChar(msg) {
		v.Pass.Reportf(call.Pos(), "log-message must be in English language")
	}
}
