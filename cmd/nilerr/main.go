package main

import (
	"github.com/golangci/nilerr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nilerr.Analyzer) }
