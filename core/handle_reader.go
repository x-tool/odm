package core

import (
	"reflect"

	"github.com/x-tool/tool"
)

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
	rawReflect *reflect.Value
	readerFieldLst
}

func newreader(i interface{}, h *Handle) *Reader {
	r := new(Reader)
	r.raw = i
	r.rawReflect = tool.GetRealReflectValue(reflect.ValueOf(i))
	return r
}

func (r *Reader) GetRaw() interface{} {
	return r.raw
}

func (r *Reader) GetRawReflect() *reflect.Value {
	return r.rawReflect
}

// get single item fields addr
func (r *Reader) GetReaderRootItemFieldsAddr() (v []reflect.Value) {
	if r.rawReflect.Kind() == reflect.Struct {
		lenR := r.rawReflect.NumField()
		for i := 0; i < lenR; i++ {
			_v := r.rawReflect.Field(i).Addr()
			v = append(v, _v)
		}
	}
	if r.rawReflect.Kind() == reflect.Slice {

	}
	return
}

func (r *Reader) AddRow(rowValues []interface{}) {
	// raws := r.getreaderRootItemFieldAddr()

}
