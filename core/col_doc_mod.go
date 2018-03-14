package core

type DocMode interface {
	Create()
	Update()
	Delete()
	Name() string
}

var DocModeMethodMap = map[handleType]string{
	insertData: "Create",
	updateData: "Update",
	deleteData: "Delete",
}

func callDocMode(h *Handle) {
	field := h.col.doc.getDocMode()
	if field != nil {
		value := field.GetValueFromRootValue(h.setValue)
		valuePtr := value.Addr()
		method := valuePtr.MethodByName(DocModeMethodMap[h.handleType])
		method.Call(nil)
	}

}
