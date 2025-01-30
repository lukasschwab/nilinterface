package main

import (
	"github.com/lukasschwab/nilinterface/pkg/nilinterface"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nilinterface.Analyzer)
}
