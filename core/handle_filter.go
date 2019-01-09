package core

import (
	"reflect"
)

type filters []filter

type handleType int

const (
	insertData handleType = iota
	updateData
	deleteData
	queryData
)

type filterCompare string

const (
	sameCompare      filterCompare = "like"
	sameLeftCompare                = "?like"
	sameRightCompare               = "like?"
	equalCompare                   = "=="
	isNullCompare                  = "isNull"
	betweenCompare                 = "between"
	inCompare                      = "in"
)

type filterJoin string

const (
	andFilter filterJoin = "and"
	orFilter             = "or"
	notFilter            = "not"
)

type filter struct {
	Handle    *Handle
	queryKind string
	queryV    *reflect.Value
	modeV     *reflect.Value
	queryLst  []filterItem
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
	andBoxs []FilterBox
	orBoxs  []FilterBox
	child   []FilterBox
}

func (r *filter) parse(s string) {

}
