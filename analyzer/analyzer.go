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
				v.checkLogCall(node)
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

func (v *validator) checkLogCall(call *ast.CallExpr) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	if x, ok := sel.X.(*ast.Ident); ok && (x.Name == "log" ||
		x.Name == "zap" || x.Name == "slog") {
		if len(call.Args) == 0 {
			return
		}
		lit, ok := call.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}
		msg := strings.Trim(lit.Value, "\"`")
		if !rules.IsLower(msg) {
			v.Pass.Reportf(call.Pos(), "log-message must be start lowercase char")
		}
		if !rules.OnlyEnglishAndWithoutSpecChar(msg) {
			v.Pass.Reportf(call.Pos(), "log-message must be in English language")
		}
	}
}
