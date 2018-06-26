package core

import (
	"errors"
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
