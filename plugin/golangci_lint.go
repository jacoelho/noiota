package main

import (
	"github.com/jacoelho/noiota/noiota"
	"golang.org/x/tools/go/analysis"
)

// New respects the plugin signature required by golangci-lint
// https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{noiota.Analyzer}, nil
}
