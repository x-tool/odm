package core

import "fmt"

type filterTree struct {
	children     []filterTree
	isConnectOr  bool
	isConnectNot bool
}

func newFilter(values ...interface{}) (f *filterTree, err error) {
	if len(values) <= 1 {
		return _, fmt.Errorf("can't")
	}
	return
}
