package noiota

import (
	"errors"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "noiota",
	Doc:      "Reports iota usage",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("analyzer is not type *inspector.Inspector")
	}

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	fn := func(node ast.Node) {
		ident, ok := node.(*ast.Ident)
		if !ok {
			return
		}

		obj, ok := pass.TypesInfo.ObjectOf(ident).(*types.Const)
		if !ok {
			return
		}

		if obj.Name() != "iota" {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     ident.Pos(),
			End:     ident.End(),
			Message: "iota usage found",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "replace iota with the value assigment",
				},
			},
		})
	}

	insp.Preorder(nodeFilter, fn)

	return nil, nil
}
