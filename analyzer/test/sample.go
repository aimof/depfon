package test

import (
	"fmt"

	"github.com/aimof/depfon/analyzer/test/testchild"
)

func sample() {
	called()
	_ = fmt.Sprint("foo")
	_ = testchild.Child(3)
}

func called() {
	return
}
