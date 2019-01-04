package analyzer

import (
	"errors"
	"go/ast"
	"go/token"
	"log"
	"sort"
	"strings"

	"github.com/aimof/depfon/domain/entity"
)

type analyzer struct {
	root      string
	entry     string
	fset      *token.FileSet
	pkgs      map[string]*ast.Package
	funcDecls map[string]*ast.FuncDecl
	tree      *entity.FunctionTree
}

// root: repository root
// entry: "analyzer:newAnalyzer"
func newAnalyzer(root, entry string) *analyzer {

	return &analyzer{
		root:      root,
		entry:     entry,
		fset:      token.NewFileSet(),
		pkgs:      nil,
		funcDecls: nil,
		tree:      nil,
	}
}
func (ana *analyzer) analyze() error {
	err := ana.newFunctionTree()
	if err != nil {
		return err
	}
	err = ana.parseAllDirs()
	if err != nil {
		return err
	}
	ana.findFunctions()
	ana.mapTree(ana.entry)
	return nil
}

func (ana *analyzer) newFunctionTree() error {
	ft, err := entity.NewFunctionTree(ana.entry)
	if err != nil {
		return err
	}
	ana.tree = ft
	return nil
}

func (ana *analyzer) mapTree(parent string) error {
	chilren, err := ana.findChildren(parent)
	if err != nil {
		return nil
	}
	for _, child := range chilren {
		err := ana.tree.AddChild(parent, child)
		if err != nil {
			return err
		}
		ana.mapTree(child)
	}
	return nil
}

func (ana *analyzer) findChildren(name string) ([]string, error) {
	children := make([]string, 0, len(ana.funcDecls))
	if fd, ok := ana.funcDecls[name]; !ok {
		return nil, errors.New("findChild: name does not exist in ana.funcDecls")
	} else {
		n := strings.Index(name, ":")
		if n != -1 {
			children = append(children, listCallingFunctions(fd, name[:n])...)
		}
	}
	return children, nil
}

func (ana *analyzer) findFunctions() {
	ana.funcDecls = make(map[string]*ast.FuncDecl, len(ana.pkgs)*10)
	for name, pkg := range ana.pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch decl.(type) {
				case *ast.FuncDecl:
					ana.funcDecls[name+":"+decl.(*ast.FuncDecl).Name.Name] = decl.(*ast.FuncDecl)
				}
			}
		}
	}
}

func listCallingFunctions(fdecl *ast.FuncDecl, pkgName string) []string {
	s := make([]string, 0, 5)

	for _, stmt := range fdecl.Body.List {
		switch stmt.(type) {
		case *ast.AssignStmt:
			for _, expr := range stmt.(*ast.AssignStmt).Rhs {
				switch expr.(type) {
				case *ast.CallExpr:
					switch fun := expr.(*ast.CallExpr).Fun; fun.(type) {
					case *ast.SelectorExpr:
						log.Println(fun)
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
