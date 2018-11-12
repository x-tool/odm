package core

import (
	"reflect"

	"github.com/x-tool/tool"
)

type Row struct {
	reader       *Reader
	fieldAddrLst []*reflect.Value
	fieldLst     structFieldLst
}

type Reader struct {
	handle     *Handle
	raw        interface{}
	rawReflect *reflect.Value
	IsComplex  bool
}

func newReader(i interface{}, h *Handle) *Reader {
	r := new(Reader)
	r.raw = i
	r.handle = h
	r.rawReflect = tool.GetRealReflectValue(reflect.ValueOf(i))
	if r.rawReflect.Kind() == reflect.Array || r.rawReflect.Kind() == reflect.Slice {
		r.IsComplex = true
	}
	return r
}

func (r *Reader) GetRaw() interface{} {
	return r.raw
}

func (r *Reader) GetRawReflect() *reflect.Value {
	return r.rawReflect
}

func (r *Reader) NewRow() (_row *Row) {
	_row.reader = r
	var item reflect.Value
	if r.IsComplex {
		item = reflect.New(r.rawReflect.Type().Elem())
	} else {
		item = *r.rawReflect
	}
	_row.raw = &item
	if item.Kind() == reflect.Struct {
		lenR := item.NumField()
		for i := 0; i < lenR; i++ {
			_v := item.Field(i)
			_row.fieldLst = append(_row.fieldLst, _v)
		}
	} else {
		_row.fieldLst = item
	}
	return
}

// get single item fields addr
func (r *Row) GetReaderRootItemFieldsAddr() (v []reflect.Value) {
	if r.reader.IsComplex {
		if r.raw.Kind() == reflect.Struct {
			lenR := r.raw.NumField()
			for i := 0; i < lenR; i++ {
				_v := r.raw.Field(i).Addr()
				r.fieldLst = append(r.fieldLst, r.getFieldFromReflectStructField(_v))
			}
		}
	}
	return
}

func (r *Row) getFieldFromReflectStructField(v *reflect.StructField) (f *structField) {
	if v.Tag == "" {
		if r.reader.handle.IsSingleCol() {
			f = r.reader.handle.GetCol().getFieldByName(v.Name)
		}
	} else {

	}
}
