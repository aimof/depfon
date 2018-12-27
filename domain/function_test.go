package domain

import "testing"

func TestNewFunctionTree(t *testing.T) {
	got := newFunctionTree()
	if got.root != nil {
		t.Error()
	}
	if len(got.all) != 0 || got.all == nil {
		t.Error()
	}
}

func TestNewFunctionAndSetRoot(t *testing.T) {
	// Test newFunctionTree
	ft := newFunctionTree()

	// Test newFunction
	rootName := "root"
	got, err := ft.newFunction(rootName)
	if err != nil {
		t.Error()
	}
	if gotFromAll, ok := ft.all[rootName]; !ok {
		t.Error()
	} else if got != gotFromAll {
		t.Error()
	}

	// TestSetRoot (use Function above)
	err = ft.setRoot(got)
	if err != nil {
		t.Error()
	}
	if ft.root != got {
		t.Error()
	}

	// Test addChild if not duplicated
	child0Name := "child0"
	err = ft.root.addChild(child0Name)
	if err != nil {
		t.Error()
	}
	child0, ok := ft.all[child0Name]
	if !ok {
		t.Error()
	}
	if child0FromRoot, ok := ft.all[child0Name]; !ok {
		t.Error()
	} else if child0 != child0FromRoot {
		t.Error()
	}
	if rootFromChild0, ok := child0.parents[rootName]; !ok {
		t.Error()
	} else if rootFromChild0 != ft.root {
		t.Error()
	}

	// Test addChild if duplicated
	child1Name := "child1"
	grandchildName := "grandchild"
	err = ft.root.addChild(child1Name)
	if err != nil {
		t.Error()
	}
	child1, ok := ft.root.children[child1Name]
	if !ok {
		t.Error()
	}
	err = child0.addChild(grandchildName)
	if err != nil {
		t.Error()
	}

	err = child1.addChild(grandchildName)
	if err != nil {
		t.Error()
	}
	grandchild0, ok := child0.children[grandchildName]
	if !ok {
		t.Error()
	}
	grandchild1, ok := child1.children[grandchildName]
	if !ok {
		t.Error()
	}
	if grandchild0 != grandchild1 {
		t.Error()
	}

	child0FromGrandchild, ok := grandchild0.parents[child0Name]
	if !ok {
		t.Error()
	}
	if child0 != child0FromGrandchild {
		t.Error()
	}

	child1FromGrandchild, ok := grandchild0.parents[child1Name]
	if !ok {
		t.Error()
	}
	if child1 != child1FromGrandchild {
		t.Error()
	}
}
