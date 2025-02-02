// nilinterface module plugin for golangci-lint modeled on
// https://github.com/golangci/example-plugin-module-linter/blob/main/example.go
package nilinterface

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

// linterPlugin implements the register.LinterPlugin interface for
// [analysis.Analyzer].
type linterPlugin struct{}

func (l linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (l linterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
