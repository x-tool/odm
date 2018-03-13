package core

func (h *Handle) GetRootValues() []*Value {
	values := h.Col.GetRootValues(h.setValue)
	return values
}
