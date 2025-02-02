package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nilinterface",
	Doc:  "check for nil passed to interface parameters",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			for i, arg := range callExpr.Args {
				if ident, ok := arg.(*ast.Ident); ok && ident.Name == "nil" {
					if sig, ok := pass.TypesInfo.TypeOf(callExpr.Fun).(*types.Signature); ok {
						if i < sig.Params().Len() && isInterface(sig.Params().At(i).Type()) {
							pass.Reportf(arg.Pos(), "nil passed to interface parameter")
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func isInterface(t types.Type) bool {
	_, ok := t.Underlying().(*types.Interface)
	return ok
}
