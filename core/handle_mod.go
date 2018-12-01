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
				// _value, _ := field.GetValueFromRootValue(&h.writter.rawValue)
				// value = *_value
			case UpdateData:
				value = field.newValue()

			}
			valuePtr := value.Addr()
			method := valuePtr.MethodByName(DocModeMethodMap[h.handleType])
			method.Call(nil)
		}
	}
}