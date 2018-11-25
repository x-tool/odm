package core

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type HandleField struct {
	reader        *Reader
	name          string
	goDepend      dependLst
	odmDepend     dependLst
	complexValues // golang struct child slice id or map key
}

type complexValue struct {
	structId int
	fieldId  int
	value    string
}

type complexValues []complexValue

func (c complexValues) getValue(id int, fieldId int) string {
	for _, v := range c {
		if id == v.structId && fieldId == v.fieldId {
			return v.value
		}
	}
	return ""
}

func newHandleField(r *Reader, f reflect.StructField) (field *HandleField, err error) {
	field = &HandleField{
		reader: r,
		name:   f.Name,
	}
	var goDepend dependLst
	var odmDepend dependLst
	var complexValues []complexValue
	// format structFieldTag
	tag := string(f.Tag)
	odmLst := strings.Split(tag, "|")
	for i, v := range odmLst {
		reg := regexp.MustCompile(`[]+`)
		ids := reg.FindStringIndex(v)
		structName := v[:ids[0]]
		fieldPath := v[:ids[1]]
		if len(ids) == 0 || (i != 0 && ids[0] != 0) {
			return field, fmt.Errorf("can't get struct name from %v's tag", f.Name)
		}
		var _struct *odmStruct
		// first split is field from col, other field from doc
		if i == 0 {
			_col, err := r.handle.getColByStr(structName)
			if err != nil {
				return field, fmt.Errorf("can't get Col from your register structs")
			}
			_struct = &_col.odmStruct
		} else {
			_struct, err = r.handle.getStructByStr(structName)
			if err != nil {
				return field, errors.New("can't get struct from your register structs")
			}
		}

		goD, odmD, complexs, err := field.formatField(_struct, fieldPath)
		if err != nil {
			return field, fmt.Errorf("can't get struct field from struct: %v", structName)
		}
		goDepend = append(goDepend, goD...)
		odmDepend = append(odmDepend, odmD...)
		complexValues = append(complexValues, complexs...)
	}
	field.goDepend = goDepend
	field.odmDepend = odmDepend
	field.complexValues = complexValues
	return field, nil
}

func (r *HandleField) formatField(o *odmStruct, s string) (goD dependLst, odmD dependLst, complexValue []complexValue, err error) {
	if s == "" {
		return goD, odmD, complexValue, nil
	}

	return
}
