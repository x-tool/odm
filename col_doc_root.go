package odm

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

	if doc.Pid == rootPid {
		value := rV.FieldByName(doc.Name)
		return &value
	} else {
		pdoc := d.getFieldById(doc.Pid)
		_value := rV.FieldByName(pdoc.Name)
		if _value.Kind() == reflect.Ptr {
			_value = reflect.Indirect(_value)
		}
		value := _value.FieldByName(doc.Name)
		return &value
	}
}
