package core

import "reflect"

func (h *Handle) GetRootValues() []*Value {
	values := h.col.GetRootValues(h.setValue)
	return values
}

func (h *Handle) setColbyValue(r *reflect.Value) {
	h.col = h.db.GetColByName(r.Type().Name())
}
