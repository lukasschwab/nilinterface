// linters modeled on https://github.com/golangci/example-plugin-module-linter/blob/main/example.go
package linters

import (
	"go/ast"
	"go/types"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("nilinterface", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return linterPlugin{}, nil
}

type linterPlugin struct{}

func (l linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func (l linterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

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
