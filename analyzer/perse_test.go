package analyzer

import (
	"testing"
)

func TestParseAllFiles(t *testing.T) {
	ana := newAnalyzer("./test", "sample")
	err := ana.parseAllFiles()
	if err != nil {
		t.Error(err)
	}
	if p, ok := ana.pkgs["test"]; !ok || p == nil {
		t.Error()
	}
	if p, ok := ana.pkgs["testchild"]; !ok || p == nil {
		t.Error()
	}

}
