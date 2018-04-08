package core

import (
	"reflect"
)

type DocMode interface {
	Create(h *Handle)
	Update(h *Handle)
	Delete(h *Handle)
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
		var value reflect.Value
		switch h.handleType {
		case insertData:
			value = *field.GetValueFromRootValue(h.getInsertValue())
		case updateData:
			value = field.newValue()

		}
		valuePtr := value.Addr()
		method := valuePtr.MethodByName(DocModeMethodMap[h.handleType])
		method.Call(nil)
	}

}
