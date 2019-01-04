package entity

import (
	"errors"
	"fmt"
)

/*
	entities
	Everyting in this file must not be exported.
*/

const (
	defaultAllMapSize      = 20
	defaultRelationMapSize = 5
)

// functionTree is a tree structure contains all functions.
type functionTree struct {
	root *function
	all  map[string]*function
}

// function is a node.
type function struct {
	name     string
	parents  map[string]*function
	children map[string]*function
	tree     *functionTree
}

func newFunctionTree() *functionTree {
	return &functionTree{
		root: nil,
		all:  make(map[string]*function, defaultAllMapSize),
	}
}

// Newfunction make function struct from functionTree. Default map is nil.
func (ft *functionTree) newFunction(name string) (*function, error) {
	if _, exists := ft.all[name]; exists {
		return nil, fmt.Errorf("newFunction: %s is already exists", name)
	}

	f := &function{
		name:     name,
		parents:  make(map[string]*function, defaultRelationMapSize),
		children: make(map[string]*function, defaultRelationMapSize),
		tree:     ft,
	}
	ft.all[name] = f
	return f, nil
}

// setRoot doesn't change root if it already exists.
func (ft *functionTree) setRoot(f *function) error {
	if ft.root != nil {
		return errors.New("setRoot: root already exists")
	}
	ft.root = f
	return nil
}

func (pf *function) addChild(name string) error {
	if cf, ok := pf.tree.all[name]; ok {
		pf.children[name] = cf
		cf.parents[pf.name] = pf
		return nil
	}
	cf, err := pf.tree.newFunction(name)
	if err != nil {
		return err
	}
	pf.children[name] = cf
	cf.parents[pf.name] = pf
	return nil
}
