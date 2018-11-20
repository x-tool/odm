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

func newReaderField(r *Reader, f reflect.StructField) (*ReaderField, error) {
	field := &ReaderField{
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
		if len(ids) == 0 || (i != 0 && ids[0] != 0) {
			return field, fmt.Errorf("can't get struct name from %v's tag", f.Name)
		}
		_struct, err := r.handle.db.getStructByName(v[:ids[1]])
		if err != nil {
			return field, errors.New("can't get struct from your register structs")
		}
		goD, odmD, complexs, err := field.formatField(_struct, v[ids[1]:])
		if err != nil {
			return field, fmt.Errorf("can't get struct field from struct: %v", v[:ids[1]])
		}
		goDepend = append(goDepend, goD...)
		odmDepend = append(odmDepend, odmD...)
		len()
	}
	return field, nil
}

func (r *ReaderField) formatField(o *odmStruct, s string) (goD dependLst, odmD dependLst, complexValue map[int]string, err error) {
	if s == "" {
		return goD, odmD, complexValue, nil
	}

	return
}
