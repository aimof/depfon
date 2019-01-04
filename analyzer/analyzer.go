package analyzer

import (
	"go/ast"
	"go/token"
	"sort"

	"github.com/aimof/depfon/domain/entity"
)

type analyzer struct {
	root  string
	entry string
	fset  *token.FileSet
	pkgs  map[string]*ast.Package
	tree  *entity.FunctionTree
}

// root: repository root
// entry: "analyzer:newAnalyzer"
func newAnalyzer(root, entry string) *analyzer {
	return &analyzer{
		root:  root,
		entry: entry,
		fset:  token.NewFileSet(),
		pkgs:  nil,
		tree:  nil,
	}
}

func (ana analyzer) newFunctionTree() error {
	ft, err := entity.NewFunctionTree(ana.entry)
	if err != nil {
		return err
	}
	ana.tree = ft
	return nil
}

func search(parsedPkgs map[string]*ast.Package) []string {
	s := make([]string, 0, 10)
	for _, decl := range parsedPkgs["test"].Files["test/sample.go"].Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			got := callingList(decl.(*ast.FuncDecl), "test")
			if len(got) != 0 {
				s = append(s, got...)
			}
		}
	}
	return s
}
func callingList(fdecl *ast.FuncDecl, pkgName string) []string {
	s := make([]string, 0, 5)

	for _, stmt := range fdecl.Body.List {
		switch stmt.(type) {
		case *ast.AssignStmt:
			for _, expr := range stmt.(*ast.AssignStmt).Rhs {
				switch expr.(type) {
				case *ast.CallExpr:
					switch fun := expr.(*ast.CallExpr).Fun; fun.(type) {
					case *ast.SelectorExpr:
						var pkg, function string
						switch x := fun.(*ast.SelectorExpr).X; x.(type) {
						case *ast.Ident:
							pkg = x.(*ast.Ident).Name
						}
						function = fun.(*ast.SelectorExpr).Sel.Name
						s = append(s, pkg+":"+function)
					}
				}
			}
		case *ast.ExprStmt:
			switch x := stmt.(*ast.ExprStmt).X; x.(type) {
			case *ast.CallExpr:
				switch fun := x.(*ast.CallExpr).Fun; fun.(type) {
				case *ast.Ident:
					s = append(s, pkgName+":"+fun.(*ast.Ident).Name)

				}

			}
		}
	}
	sort.Strings(s)
	return s
}
