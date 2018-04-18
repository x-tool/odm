package core

import (
	"reflect"
)

type query struct {
	Handle    *Handle
	queryKind string
	queryV    *reflect.Value
	modeV     *reflect.Value
	queryLst  []queryItem
	limitNum  int
	limitDesc bool
}

type queryItem struct {
	field      structField
	whereCheck string
	whereAnd   bool
}

func newQuery(o *Handle) *query {
	r := &query{
		Handle: o,
		// queryV:    rV,
		// queryKind: t,
	}
	r.setDependToDoc()
	return r
}
func newqueryWithoutCol(rV *reflect.Value) *query {
	r := &query{
		queryV: rV,
	}
	return r
}

func (r *query) setDependToDoc() {
	// T := r.queryV.Type()
	// var value reflect.Type
	// if T.Kind() == reflect.Slice {
	// 	value = T.Elem()
	// } else {
	// 	value = T
	// }
	// var valueItem reflect.Value
	// var valueItemT reflect.Type
	// if value.Kind() == reflect.Slice {
	// 	valueItem = r.queryV.Elem()
	// } else {
	// 	valueItem = *r.queryV
	// }
	// valueItemT = valueItem.Type()
	// for i := 0; i < valueItem.NumField(); i++ {
	// 	field := valueItem.Field(i)
	// 	fieldT := valueItemT.Field(i)
	// 	if isDocMode(fieldT.Name) {
	// 		r.modeV = &field
	// 	}
	// 	newqueryItem := r.Handle.DependToDoc(strings.Split(fieldT.Tag.Get(tagName), "."), fieldT.Name)
	// 	r.queryLst = append(r.queryLst, newqueryItem)
	// }

}
