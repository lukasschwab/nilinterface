package analyzer_test

import (
	"testing"

	"github.com/lukasschwab/nilinterface/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "./...")
}
