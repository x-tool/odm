package core

import (
	"reflect"
	"strings"
)

const (
	splitStructStr = ":"
)

func (d *doc) findDocModeField() (field *structField) {
	for _, v := range d.getExtendFields() {
		_value := reflect.New(v.sourceType)
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

// "@tag"
// "path"
// "path:@tag"
// "path:path"
func (d *doc) getFieldValue(s string) (f *structField, err error) {
	structLst := strings.Split(s, splitStructStr)
	if len(structLst) == 1 {
		return d.getFieldByString()
	} else {
		d.getFieldByString(structLst[0])

	}
}
