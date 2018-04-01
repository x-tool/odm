package core

import (
	"encoding/json"
	"reflect"
	"strings"
)

/// error
func (d *structField) GetValueFromRootValue(rootValue *reflect.Value) *reflect.Value {
	_r := *rootValue
	if len(d.dependLst) != 0 {
		for _, v := range d.dependLst {
			// if dependLst parent is extend, should get field from grandparent
			if v.kind == Struct {
				_r = _r.FieldByName(v.Name())
			} else {
				break // !!!!!! mark wait modify !!!!!!!
			}
		}
		_r = _r.FieldByName(d.Name())
	} else {
		_r = _r.FieldByName(d.Name())
	}

	return &_r
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}

func (d *structField) getFieldByStr(s string) (f *structField) {
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

func getFieldByDependLst(fields structFieldLst, Lst []string) (d *structField) {
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
