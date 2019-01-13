package core

import (
	"reflect"
)

var formatBracketStr string

type filters []filter

const (
	andLink linkKind = iota
	orLink
	notLink
)

type filter struct {
	Handle    *Handle
	queryKind string
	queryV    *reflect.Value
	modeV     *reflect.Value
	limitNum  int
	limitDesc bool
}

func newFilter(o *Handle) *filter {
	r := &filter{
		Handle: o,
		// filterV:    rV,
		// filterKind: t,
	}
	// r.setDependToDoc()
	return r
}

func (f *filter) parse(s string, values ...interface{}) (rootBox ASTTree, err error) {
	rootBox = ASTTree{source: s}
	err = setBracketsTree(&rootBox)
	return rootBox, err
}
