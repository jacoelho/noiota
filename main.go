package main

import (
	"github.com/jacoelho/noiota/noiota"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(noiota.Analyzer)
}
