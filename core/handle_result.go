package core

import "reflect"

type result struct {
	handle     *Handle
	raw        interface{}
	rawReflect reflect.Value
	resultFieldLst
}

type resultFieldLst []*resultField
type resultField struct {
	field    *structField
	function resultFunction
}

type resultFunction int

const (
	resultNumFunction resultFunction = iota
)

func newResult(i interface{}) *result {
	reflect.TypeOf(i)
	r := new(result)
	r.raw = i

	return r
}

// type result struct {
// 	Col            *Col
// 	resultFieldLst []*docField
// 	resultV        *reflect.Value
// 	resultKind     int
// 	resultElem     *reflect.Value
// }

// func newResult(rV *reflect.Value, c *Col) (r *result) {
// 	var vK int
// 	var vE reflect.Value
// 	if rV.Kind() == reflect.Slice {
// 		vK = 0
// 		vE = rV.Elem()
// 	} else {
// 		vK = 1
// 		vE = *rV
// 	}
// 	r = &result{
// 		Col:        c,
// 		resultV:    rV,
// 		resultKind: vK,
// 		resultElem: &vE,
// 	}
// 	return
// }
// func (r *result) newResultItem() (v *reflect.Value) {
// 	var rV reflect.Value
// 	if r.resultKind == 0 {
// 		rV = reflect.New(r.resultElem.Type())
// 	} else {
// 		rV = reflect.New(r.resultV.Type())
// 	}
// 	return &rV
// }

func (r *result) getResultRootItemFieldAddr(rootV *reflect.Value) (v []reflect.Value) {
	if rootV.Kind() == reflect.Struct {
		lenR := rootV.NumField()
		for i := 0; i < lenR; i++ {
			_v := rootV.Field(i).Addr()
			v = append(v, _v)
		}
	}
	return
}

func (r *result) AddRow(rowValues []interface{}) {
	raws := r.getResultRootItemFieldAddr()

}
