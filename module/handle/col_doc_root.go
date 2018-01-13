package core

import "reflect"

func (doc *Doc) getRootExtendFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.Pid == -1 && v.extendPid != -1 && v.isExtend {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootSinpleFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid == -1 && !v.isExtend && !doc.checkComplexField(v) {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootComplexFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid == -1 && !v.isExtend && doc.checkComplexField(v) {
			d = append(d, v)
		}
	}
	return
}

func (d *Doc) getRootDetails() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !v.isExtend {
			doc = append(doc, v)
		}
	}
	return
}
func (d *Doc) getRootDetailsWithExtend() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 {
			doc = append(doc, v)
		}
	}
	return
}

func (d *Doc) getRootDetailValue(rootValue *reflect.Value, doc *DocField) (v *reflect.Value) {
	rV := *rootValue
	if rV.Kind() == reflect.Ptr {
		rV = reflect.Indirect(rV)
		if rV.Kind() != reflect.Struct {
			return nil
		}
	}

	if doc.Pid == -1 {
		value := rV.FieldByName(doc.Name)
		return &value
	} else {
		pDoc := d.getFieldById(doc.Pid)
		_value := rV.FieldByName(pDoc.Name)
		if _value.Kind() == reflect.Ptr {
			_value = reflect.Indirect(_value)
		}
		value := _value.FieldByName(doc.Name)
		return &value
	}
}
