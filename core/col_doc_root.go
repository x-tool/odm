package core

import "reflect"

const (
	rootPid = -1
)

func (d *doc) getRootDetails() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !v.isExtend {
			doc = append(doc, v)
		}
	}
	return
}
func (d *doc) getRootDetailsWithExtend() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 {
			doc = append(doc, v)
		}
	}
	return
}

func (d *doc) getRootDetailValue(rootValue *reflect.Value, doc *docField) (v *reflect.Value) {
	rV := *rootValue
	if rV.Kind() == reflect.Ptr {
		rV = reflect.Indirect(rV)
		if rV.Kind() != reflect.Struct {
			return nil
		}
	}

	if doc.pid == rootPid {
		value := rV.FieldByName(doc.GetName())
		return &value
	} else {
		pdoc := d.getFieldById(doc.pid)
		_value := rV.FieldByName(pdoc.GetName())
		if _value.Kind() == reflect.Ptr {
			_value = reflect.Indirect(_value)
		}
		value := _value.FieldByName(doc.GetName())
		return &value
	}
}
