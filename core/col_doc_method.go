package core

import (
	"errors"
	"reflect"
	"strings"
)

func (d *doc) findDocModeField() (field *structField) {
	for _, v := range d.getExtendFields() {
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

// "path:structName.path..." // if doc field is interface, use ":" to get odmStruct from database, and search path
func (d *doc) getFieldByDocPath(path string) (f *structField, err error) {
	// check pathLst
	pathLst := strings.Split(path, ":")
	pathLen := len(pathLst)
	if pathLen == 1 {
		f = d.getFieldByPath(pathLst[0])
	} else {
		var acrossDependLst dependLst
		f = d.getFieldByPath(pathLst[0])
		acrossDependLst = append(acrossDependLst, f.dependLst...)

		for _, v := range pathLst {
			// add fieldSelf to acrossDependLst
			acrossDependLst = append(acrossDependLst, f)
			structPath := strings.Split(v, ".")
			structName := structPath[0]
			_odmStruct := d.col.database.getStructByName(structName)
			f = _odmStruct.getFieldByPath(strings.Join(structPath[1:], "."))
			if f == nil {
				err = errors.New("canot get field by path '" + v + "'")
				return
			} else {
				acrossDependLst = append(acrossDependLst, f.dependLst...)
			}
		}
		f.AcrossStructdependLst = acrossDependLst
	}
	return
}
