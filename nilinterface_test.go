package linters_test

import (
	"testing"

	linters "github.com/lukasschwab/nilinterface"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, linters.Analyzer, "a")
	analysistest.Run(t, testdata, linters.Analyzer, "b")
	analysistest.Run(t, testdata, linters.Analyzer, "c")
}
