package core

import "reflect"

type writter struct {
	handle   *Handle
	setLst   []*writeItem
	raw      interface{}
	rawValue reflect.Value
}

type writeItem struct {
	HandleFieldLst
	value interface{}
}

func newWritter(h *Handle) *writter {
	w := &writter{
		handle: h,
	}
	return w
}

func (w *writter) setWritterValue(i interface{}) {
	w.raw = i
	w.rawValue = reflect.ValueOf(i)
}

func (w *writter) GetWritterValue() *reflect.Value {
	return &w.rawValue
}
func (w *writter) add(d dependLst, value interface{}) {
	item := &writeItem{
		dependLstToRoot: d,
		value:           value,
	}
	w.setLst = append(w.setLst, item)
}
