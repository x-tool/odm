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
	InsertData: "Create",
	UpdateData: "Update",
	DeleteData: "Delete",
}

func callDocMode(h *Handle) {
	for _, v := range h.handleCols {
		field := v.col.getDocMode()
		if field != nil {
			var value reflect.Value
			switch h.handleType {
			case InsertData:
				value = field.newValue()
			case UpdateData:
				value = field.newValue()

			}
			method := value.MethodByName(DocModeMethodMap[h.handleType])
			method.Call(make([]reflect.Value, 0))
		}
	}
}
