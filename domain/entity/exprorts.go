package entity

import (
	"errors"
	"sort"
)

// FunctionTree is a tree structure of funcitons.
type FunctionTree struct {
	tree *functionTree
}

// NewFunctionTree Initializes FucntionTree and set name as root.
func NewFunctionTree(name string) (*FunctionTree, error) {
	ft := newFunctionTree()
	f, err := ft.newFunction(name)
	if err != nil {
		return nil, err
	}
	err = ft.setRoot(f)
	if err != nil {
		return nil, err
	}
	return &FunctionTree{
		tree: ft,
	}, nil
}

// Show shows name's dependency
func (ft *FunctionTree) Show(name string) (parents, children []string, err error) {
	f, ok := ft.tree.all[name]
	if !ok {
		return nil, nil, errors.New("Show: name does not exist")
	}
	parentsSlice := make([]string, 0, len(f.parents))
	for key := range f.parents {
		parentsSlice = append(parentsSlice, key)
	}
	sort.Strings(parentsSlice)
	childrenSlice := make([]string, 0, len(f.children))
	for key := range f.children {
		childrenSlice = append(childrenSlice, key)
	}
	sort.Strings(childrenSlice)
	return parentsSlice, childrenSlice, nil
}

// CountAll counts all functions in FunctionTree
func (ft *FunctionTree) CountAll() int {
	return len(ft.tree.all)
}

// Has returns FunctionTree has name or not.
func (ft *FunctionTree) Has(name string) bool {
	_, ok := ft.tree.all[name]
	return ok
}

// AddChild makes relation parent from child.
// parent needs to exist in FunctionTree.
func (ft *FunctionTree) AddChild(parentName, childName string) error {
	parent, ok := ft.tree.all[parentName]
	if !ok {
		return errors.New("AddChild: parent doesn't exist")
	}
	return parent.addChild(childName)
}
