package domain_test

import (
	"testing"

	"github.com/aimof/depfon/domain"
)

var (
	rootName       = "root"
	child0Name     = "child0"
	child1Name     = "child1"
	grandchildName = "grandchildName"
	dummyName      = "dummy"
)

// make sampleFunctionTree and TestNewFunctionTree
func sampleFunctionTree(t *testing.T) *domain.FunctionTree {

	ft, err := domain.NewFunctionTree(rootName)
	if ft == nil {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
	err = ft.AddChild(rootName, child0Name)
	if err != nil {
		t.Error()
	}
	err = ft.AddChild(rootName, child1Name)
	if err != nil {
		t.Error()
	}
	err = ft.AddChild(child0Name, grandchildName)
	if err != nil {
		t.Error()
	}
	err = ft.AddChild(child1Name, grandchildName)
	if err != nil {
		t.Error()
	}
	return ft
}

func TestShow(t *testing.T) {
	sft := sampleFunctionTree(t)

	// test child0
	p, c, err := sft.Show(child0Name)
	if err != nil {
		t.Error()
	}
	if len(p) != 1 || p[0] != rootName {
		t.Error()
	}
	if len(c) != 1 || c[0] != grandchildName {
		t.Error()
	}

	// test grandchild
	p, c, err = sft.Show(grandchildName)
	if err != nil {
		t.Error()
	}
	if len(p) != 2 || p[0] != child0Name || p[1] != child1Name {
		t.Error()
	}
	if len(c) != 0 {
		t.Error()
	}
}

func TestCountAll(t *testing.T) {
	ft, err := domain.NewFunctionTree(rootName)
	if err != nil {
		t.Error()
	}
	if ft.CountAll() != 1 {
		t.Error()
	}
	sft := sampleFunctionTree(t)
	if sft.CountAll() != 4 {
		t.Error()
	}
}

func TestHas(t *testing.T) {
	sft := sampleFunctionTree(t)
	if !sft.Has(child1Name) {
		t.Error()
	}
	if sft.Has(dummyName) {
		t.Error()
	}
}
