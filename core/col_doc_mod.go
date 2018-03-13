package core

type DocMode interface {
	Create()
	Update()
	Delete()
	Name()
}

var DocModeMethodMap = map[handleType]string{
	insertData: "Create",
	updateData: "Update",
	deleteData: "Delete",
}

func callDocMode(h *Handle) {
	field := h.Col.doc.getDocMode()
	if field != nil {
		value := field.GetValueFromRootValue(h.setValue)
		method := value.MethodByName(DocModeMethodMap[h.handleType])
		method.Call(nil)
	}

}
