package core

import "reflect"

type reader struct {
	handle     *Handle
	raw        interface{}
	rawReflect reflect.Value
	readerFieldLst
}

type readerFieldLst []*readerField
type readerField struct {
	field    *structField
	function readerFunction
}

type readerFunction int

const (
	readerNumFunction readerFunction = iota
)

func newreader(i interface{}) *reader {
	reflect.TypeOf(i)
	r := new(reader)
	r.raw = i

	return r
}

// type reader struct {
// 	Col            *Col
// 	readerFieldLst []*docField
// 	readerV        *reflect.Value
// 	readerKind     int
// 	readerElem     *reflect.Value
// }

// func newreader(rV *reflect.Value, c *Col) (r *reader) {
// 	var vK int
// 	var vE reflect.Value
// 	if rV.Kind() == reflect.Slice {
// 		vK = 0
// 		vE = rV.Elem()
// 	} else {
// 		vK = 1
// 		vE = *rV
// 	}
// 	r = &reader{
// 		Col:        c,
// 		readerV:    rV,
// 		readerKind: vK,
// 		readerElem: &vE,
// 	}
// 	return
// }
// func (r *reader) newreaderItem() (v *reflect.Value) {
// 	var rV reflect.Value
// 	if r.readerKind == 0 {
// 		rV = reflect.New(r.readerElem.Type())
// 	} else {
// 		rV = reflect.New(r.readerV.Type())
// 	}
// 	return &rV
// }

func (r *reader) getreaderRootItemFieldAddr(rootV *reflect.Value) (v []reflect.Value) {
	if rootV.Kind() == reflect.Struct {
		lenR := rootV.NumField()
		for i := 0; i < lenR; i++ {
			_v := rootV.Field(i).Addr()
			v = append(v, _v)
		}
	}
	return
}

func (r *reader) AddRow(rowValues []interface{}) {
	raws := r.getreaderRootItemFieldAddr()

}
