package core

import (
	"reflect"
)

type filters []filter

type CompareKind int

const (
	likeCompare    CompareKind = iota
	equalCompare               // ==
	isNullCompare              // isNull
	betweenCompare             // between
	inCompare                  // in
)

type linkKind int

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

type FilterBox struct {
	linkKind
	child []FilterBox
	field *StructField
	CompareKind
	valueLst []reflect.Value
}

func (f *filter) parse(s string) (box FilterBox) {
	return
}
