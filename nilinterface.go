// linters modeled on https://github.com/golangci/example-plugin-module-linter/blob/main/example.go
package linters

import (
	"github.com/lukasschwab/nilinterface/pkg/analyzer"

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
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (l linterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
