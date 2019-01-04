package testchild

import (
	"strconv"

	grandchild "github.com/aimof/depfon/analyzer/test/testchild/testgrandchild"
)

func Child(n int) string {
	err := grandchild.Grandchild()
	if err != nil {
		return ""
	}
	return strconv.Itoa(n)
}
