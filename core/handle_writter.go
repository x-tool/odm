package core

import (
	"fmt"
	"reflect"

	"github.com/x-tool/tool"
)

type writter struct {
	handle    *Handle
	setLst    []*writeItem // use for update or insert some filed
	raw       interface{}
	rawValue  reflect.Value
	isComplex bool
}

type writeItem struct {
	handleField *HandleField
	value       interface{}
}

func newWritter(h *Handle) *writter {
	w := &writter{
		handle: h,
	}
	return w
}

//////////// for insert value

// insert value method
func (w *writter) setWritterValue(i interface{}) {
	w.raw = i
	w.rawValue = tool.GetRealReflectValue(reflect.ValueOf(i))
	w.isComplex = tool.IsComplex(w.rawValue)
	// get col name
	item := w.rawValue
	if w.isComplex {
		item = w.rawValue.Index(0)
	}
	itemType := item.Type()
	col, err := w.handle.db.GetCol(itemType.Name())
	if err != nil {
		w.handle.Err = err
	}
	_handleCol := newHandleCol(col)
	w.handle.handleCols.add(_handleCol)
}

// get insert value
func (w *writter) GetWritterValue() *reflect.Value {
	return &w.rawValue
}

func (w *writter) IsComplexWritter() bool {
	return w.isComplex
}

/////////// for update value

// add write field
func (w *writter) add(f reflect.StructField, value interface{}) error {
	field, err := newHandleField(w.handle, string(f.Tag))
	if err != nil {
		return fmt.Errorf("can't get Field: %v use tag %v", f.Name, f.Tag)
	}
	item := &writeItem{
		handleField: field,
		value:       value,
	}
	w.setLst = append(w.setLst, item)
	return nil
}
