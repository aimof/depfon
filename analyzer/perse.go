package analyzer

import (
	"errors"
	"go/ast"
	"go/parser"
	"io/ioutil"
	"os"
)

func (ana *analyzer) parseAllDirs() error {
	pkgs, err := ana.digPkgs(ana.root)
	if err != nil {
		return err
	}
	ana.pkgs = pkgs
	return nil
}

func (ana *analyzer) digPkgs(dirName string) (map[string]*ast.Package, error) {
	fi, err := os.Stat(dirName)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("dirName is not a directory")
	}

	pkgs, err := parser.ParseDir(ana.fset, dirName, nil, 0)
	if err != nil {

	}
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			childPkgs, err := ana.digPkgs(dirName + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			pkgs = merge(childPkgs, pkgs)
		}
	}
	return pkgs, nil
}

func merge(map0, map1 map[string]*ast.Package) map[string]*ast.Package {
	joined := make(map[string]*ast.Package, len(map0)+len(map1))
	for k, v := range map0 {
		joined[k] = v
	}
	for k, v := range map1 {
		joined[k] = v
	}
	return joined
}
