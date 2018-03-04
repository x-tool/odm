package core

// func (d *doc) GetRootDetailValue(rootValue *reflect.Value, field *docField) (v *reflect.Value) {
// 	rV := *rootValue
// 	if rV.Kind() == reflect.Ptr {
// 		rV = reflect.Indirect(rV)
// 		if rV.Kind() != reflect.Struct {
// 			return nil
// 		}
// 	}

// 	if doc.parent == nil {
// 		value := rV.FieldByName(doc.GetName())
// 		return &value
// 	} else {
// 		pdoc := d.getFieldById(doc.pid)
// 		_value := rV.FieldByName(pdoc.GetName())
// 		if _value.Kind() == reflect.Ptr {
// 			_value = reflect.Indirect(_value)
// 		}
// 		value := _value.FieldByName(doc.GetName())
// 		return &value
// 	}
// }
