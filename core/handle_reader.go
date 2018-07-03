package core

import "reflect"

type reader struct {
	handle     *Handle
	isSingle   bool // true read single item ,false read item slice
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
