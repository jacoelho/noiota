package noiota_test

import (
	"testing"

	"github.com/jacoelho/noiota/noiota"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), noiota.Analyzer)
}
