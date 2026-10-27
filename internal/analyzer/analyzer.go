package analyzer

import (
	"fmt"
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

		fmt.Println(funcTarget.FullName())
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
