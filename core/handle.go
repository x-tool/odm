package core

import (
	"context"
	"reflect"
)

type handleType int

const (
	insertData handleType = iota
	updateData
	deleteData
	queryData
)

type filterCompare string

const (
	sameCompare      filterCompare = "like"  // like
	sameCompareLeft                = "?like" // ??like
	sameCompareRight               = "like?" // like??
	equalCompare                   = "=="    // ==
	isNullCompare                  = "isNull"
	betweenCompare                 = "between"
	inCompare                      = "in"
)

type filterConnect string

const (
	andFilter filterConnect = "and"
	orFilter                = "or"
	notFilter               = "not"
)

type filters []filter

type filter struct {
	target       *structField
	compare      filterCompare
	value        interface{}
	connect      filterConnect
	childFilters filters
}

type setValueLst []*setValue
type setValue struct {
	value *reflect.Value
	filter
}

func newSetValue(value *reflect.Value, f filter) (s *setValue) {
	s = &setValue{
		value:  value,
		filter: f,
	}
	return
}

type Handle struct {
	handleType
	setValueLst
	filterDoc filter
	db        *Database
	col       *Col
	context   context.Context
	result    interface{}

	Err error
}

func (d *Handle) GetDBName() string {
	return d.col.database.name
}

func (d *Handle) GetColName() string {
	return d.col.Name()
}

func (d *Handle) GetCol() *Col {
	return d.col
}

func (d *Handle) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}

// func (h *Handle) GetRootValues() []*Value {
// 	values := h.col.GetRootValues(h.setValue)
// 	return values
// }

func (h *Handle) setColbyValue(r *reflect.Value) {
	if h.col != nil {
		return
	}
	h.col = h.db.GetColByName(r.Type().Name())
}

func (h *Handle) addSetValue(s *setValue) {
	h.setValueLst = append(h.setValueLst, s)
}

func (h *Handle) getInsertValue() *reflect.Value {
	return h.setValueLst[0].value
}

func newHandle(db *Database, con context.Context) *Handle {
	d := &Handle{
		db:      db,
		context: con,
	}
	return d

}

func newHandleByCol(c *Col, con context.Context) *Handle {
	d := &Handle{
		db:      c.database,
		col:     c,
		context: con,
	}
	return d

}
