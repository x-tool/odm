package core

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type ReaderField struct {
	reader       *Reader
	name         string
	goDepend     dependLst
	odmDepend    dependLst
	complexValue map[int]string // golang struct child slice id or map key
}

func newReaderField(r *Reader, f reflect.StructField) (field *ReaderField, err error) {
	field = &ReaderField{
		reader: r,
		name:   f.Name,
	}
	var goDepend dependLst
	var odmDepend dependLst
	var complexValue map[int]string
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
		complexValue = append(complexValue, complexs...)
	}
	field.goDepend = goDepend
	field.odmDepend = odmDepend
	field.complexValue = complexValue
	return field, nil
}

func (r *ReaderField) formatField(o *odmStruct, s string) (goD dependLst, odmD dependLst, complexValue map[int]string, err error) {
	if s == "" {
		return goD, odmD, complexValue, nil
	}

	return
}
