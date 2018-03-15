package core

import (
	"reflect"
	"strings"
)

func (d *doc) findDocMode() (field *docField) {
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

func (d *doc) getDocMode() *docField {
	return d.mode
}

func (d *doc) getDocFieldByStr(s string) (f *docField) {
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

func getFieldByDependLst(fields docFieldLst, Lst []string) (d *docField) {
	for _, field := range fields {
		var check bool = false
		for i, dependField := range field.dependLst {
			if dependField.Name() != Lst[i] {
				check = false
				break
			} else {
				check = true
			}
		}
		if check == true {
			d = field
			break
		}
	}
	return
}
