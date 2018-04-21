package core

import (
	"reflect"
	"strings"
)

func (d *doc) findDocMode() (field *structField) {
	for _, v := range d.getStructRootFields() {
		_value := reflect.New(v.selfType)
		_, ok := _value.Interface().(DocMode)
		if ok {
			field = v
			break
		}
	}
	return
}

func (d *doc) getDocMode() *structField {
	return d.mode
}

// "fieldName"
// "tagName"
// "path.fieldName"
func (d *doc) getDocFieldByStr(s string) (f *structField) {
	// check dependLst
	var dependLst []string
	dependLst = strings.Split(s, ".")
	dependLen := len(dependLst)
	// if has no dependLst Or is root field
	if dependLen == 1 {
		// get tag first
		byTag := d.getFieldByTag(s)
		if byTag != nil {
			f = byTag
		} else {
			fLst := d.getFieldByName(s)
			fLstLen := len(fLst)
			// if docFieldLstLen != 1 return nil
			if fLstLen == 1 {
				f = fLst[0]
			} else {
				f = nil
			}
		}
	} else {
		fields := d.getFieldByName(dependLst[dependLen-1])
		// if docFieldLstLen != 1 range depend
		if len(fields) == 1 {
			f = fields[0]
		} else {
			f = getFieldByDependLst(fields, dependLst)
		}
	}
	return
}
