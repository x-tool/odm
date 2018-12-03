package core

import (
	"fmt"
	"reflect"

	"github.com/x-tool/tool"
)

type writter struct {
	handle    *Handle
	setLst    []*writeItem
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
	_handleCol := newHandleCol(w.handle.db.GetCol(itemType.Name()))
	w.handle.handleCols.add(_handleCol)
}

func (w *writter) GetWritterValue() *reflect.Value {
	return &w.rawValue
}
func (w *writter) IsComplex() bool {
	return w.isComplex
}

func (w *writter) add(f reflect.StructField, value interface{}) error {
	field, err := newHandleField(w.handle, f)
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
