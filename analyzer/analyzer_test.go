package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestFindFunctions(t *testing.T) {
	ana := newAnalyzer("./test", "")
	ana.pkgs = parseTestPkg(t)
	ana.findFunctions()
	if _, ok := ana.funcDecls["test:sample"]; !ok {
		t.Error()
	}
}
func TestListCallingFuncions(t *testing.T) {
	parsedPkgs := parseTestPkg(t)
	want := []string{"fmt:Sprint", "test:called", "testchild:Child"}
	got := make([]string, 0, 10)
	for _, decl := range parsedPkgs["test"].Files["test/sample.go"].Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			funcs := listCallingFunctions(decl.(*ast.FuncDecl), "test")
			got = append(got, funcs...)
		}
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}

func parseTestPkg(t *testing.T) map[string]*ast.Package {
	fset := token.NewFileSet()
	parsedPkgs, err := parser.ParseDir(fset, "./test", nil, 0)
	if err != nil {
		t.Error()
	}
	return parsedPkgs
}
