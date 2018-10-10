package core

import "reflect"

// return col item type, Ex: by fieldname, by count, by list...
type readerType int

const (
	readField readerType = iota
	readFunc
)

// return col list
type readerFieldLst []*readerField

type readerField struct {
	readerType
	dependLst
	function readerFunction
}

type readerFunction int

const (
	readerNumFunction readerFunction = iota
)

type Reader struct {
	raw        interface{}
	rawReflect reflect.Value
	readerFieldLst
}

func newreader(i interface{}) *Reader {
	reflect.TypeOf(i)
	r := new(Reader)
	r.raw = i
	return r
}

func (r *Reader) getreaderRootItemFieldAddr(rootV *reflect.Value) (v []reflect.Value) {
	if rootV.Kind() == reflect.Struct {
		lenR := rootV.NumField()
		for i := 0; i < lenR; i++ {
			_v := rootV.Field(i).Addr()
			v = append(v, _v)
		}
	}
	return
}

func (r *Reader) AddRow(rowValues []interface{}) {
	// raws := r.getreaderRootItemFieldAddr()

}
