package core

func (h *Handle) GetRootValues() []*Value {
	values := h.col.GetRootValues(h.setValue)
	return values
}
