package analyzer

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestSearchFunctionList(t *testing.T) {
	fset := token.NewFileSet()
	parsedPkgs, err := parser.ParseDir(fset, "./test", nil, 0)
	if err != nil {
		t.Error()
	}
	want := []string{"fmt:Sprint", "test:called", "testchild:Child"}
	got := search(parsedPkgs)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}
