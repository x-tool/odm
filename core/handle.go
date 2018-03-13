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
)

type contrast int

const (
	sameData      contrast = iota // like
	sameDataLeft                  // ??like
	sameDataRight                 // like??
	equalData                     // ==
)

type HandleFilter struct {
	target   *docField
	contrast contrast
	value    interface{}
}

func (h HandleFilter) Kind() Kind {
	return h.target.GetKind()
}
func (h HandleFilter) FieldName() string {
	return h.target.GetName()
}
func (h HandleFilter) Vakue() interface{} {
	return h.value
}

type HandleFilterLst []*HandleFilter

type HandleSetValue struct {
	target     *docField
	handleType handleType
	value      interface{}
}

func (h HandleSetValue) Kind() Kind {
	return h.target.GetKind()
}
func (h HandleSetValue) FieldName() string {
	return h.target.GetName()
}
func (h HandleSetValue) Vakue() interface{} {
	return h.value
}

type HandleSetValueLst []*HandleSetValue

type HandleGroup struct {
	filterLst HandleFilterLst
	setLst    HandleSetValue
}

type Handle struct {
	// ptr to Col
	Col *Col
	handleType
	context        context.Context
	filterDocs     HandleFilterLst
	HandleGroupLst []*HandleGroup
	result         interface{}
	setValue       *reflect.Value

	Err error
}

func (d *Handle) GetDBName() string {
	return d.Col.database.name
}

func (d *Handle) GetColName() string {
	return d.Col.GetName()
}

func (d *Handle) Value() *reflect.Value {
	return d.setValue
}

func (d *Handle) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}

func newHandle(col *Col, h handleType, con context.Context) *Handle {
	d := &Handle{
		Col:        col,
		handleType: h,
		context:    con,
	}
	return d

}
