package main

import (
	"github.com/lukasschwab/nilinterface/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

// main is an entry-point for running the nilinterface analyzer as a standalone
// binary. The analyzer process exposed by [singlechecker.Main] is additionally
// compatible with `go vet -vettool`.
func main() {
	singlechecker.Main(analyzer.Analyzer)
}
