package core

import (
	"errors"
	"reflect"
)

type Row struct {
	reader       *Reader
	raw          reflect.Value
	fieldAddrLst []reflect.Value
}

type Reader struct {
	handle     *Handle
	raw        interface{}
	rawReflect reflect.Value
	itemLst    []*ReaderField // result item one row field list
	IsComplex  bool
}

// result item type should be ptr
func newReader(i interface{}, h *Handle) (*Reader, error) {
	r := new(Reader)
	r.raw = i
	r.handle = h
	// if raw type is not ptr, can't white result in it, return error
	rawValue := reflect.ValueOf(i)
	if r.rawReflect.Kind() != reflect.Ptr {
		return r, errors.New("Result type should be Ptr")
	}
	r.rawReflect = reflect.Indirect(rawValue)
	if r.rawReflect.Kind() == reflect.Array || r.rawReflect.Kind() == reflect.Slice {
		r.IsComplex = true
	}
	return r, nil
}

func (r *ReaderField) formatFields(s string) {
	r.rawReflect
}

// if result raw value is complex type return new row
// if is single return raw
func (r *Reader) Row() (_row *Row) {
	_row.reader = r
	var item reflect.Value
	if r.IsComplex {
		item = reflect.New(r.rawReflect.Type().Elem())
		_raw := reflect.Indirect(r.rawReflect)
		if _raw.CanSet() {
			reflect.Append(_raw, item)
		}
	} else {
		item = r.rawReflect
	}
	_row.raw = item
	// push field ptr
	if item.Kind() == reflect.Struct {
		lenR := item.NumField()
		for i := 0; i < lenR; i++ {
			_v := item.Field(i)
			_row.fieldAddrLst = append(_row.fieldAddrLst, _v.Addr())
		}
	} else {
		_row.fieldAddrLst = append(_row.fieldAddrLst, item.Addr())
	}
	return
}
