package analyzer_test

import (
	"testing"

	"github.com/lukasschwab/nilinterface/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "a")
	analysistest.Run(t, testdata, analyzer.Analyzer, "b")
	analysistest.Run(t, testdata, analyzer.Analyzer, "c")
	analysistest.Run(t, testdata, analyzer.Analyzer, "d")
	analysistest.Run(t, testdata, analyzer.Analyzer, "e")
	analysistest.Run(t, testdata, analyzer.Analyzer, "f")
}
