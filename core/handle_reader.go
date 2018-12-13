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
	handle       *Handle
	raw          interface{}
	rawReflect   reflect.Value
	itemFieldLst []*HandleField // result item one row field list
	isComplex    bool
}

// result item type should be ptr
// result item raw type should be struct
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
		r.isComplex = true
	}
	if r.rawReflect.Elem().Kind() != reflect.Struct {
		return r, errors.New("Result indirect type should be Struct")
	}
	err := r.formatFields()
	return r, err
}

func (r *Reader) IsComplex() bool {
	return r.isComplex
}

func (r *Reader) formatFields() error {
	var itemType reflect.Type
	if r.isComplex {
		itemType = r.rawReflect.Elem().Type()
	} else {
		itemType = r.rawReflect.Type()
	}
	itemFieldLen := itemType.NumField()
	for i := 0; i < itemFieldLen; i++ {
		fieldStruct := itemType.Field(i)
		readField, err := newHandleField(r.handle, string(fieldStruct.Tag))
		if err != nil {
			return err
		}
		r.itemFieldLst = append(r.itemFieldLst, readField)
	}
	return nil
}

// if result raw value is complex type return new row
// if is single return raw
func (r *Reader) Row() (_row *Row) {
	_row.reader = r
	var item reflect.Value
	if r.isComplex {
		item = reflect.New(r.rawReflect.Elem().Type())
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

func (r *Row) Addrs() (result []interface{}) {
	for _, v := range r.fieldAddrLst {
		result = append(result, v.Interface())
	}
	return
}
