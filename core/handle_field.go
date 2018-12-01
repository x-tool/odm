package core

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type HandleField struct {
	handle        *Handle
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

func newHandleField(h *Handle, f reflect.StructField) (field *HandleField, err error) {
	field = &HandleField{
		handle: h,
		name:   f.Name,
	}
	var goDepend dependLst
	var odmDepend dependLst
	var complexValues []complexValue
	// format structFieldTag
	tag := string(f.Tag)
	odmLst := strings.Split(tag, "|")
	for i, v := range odmLst {
		reg := regexp.MustCompile(`\w`)
		ids := reg.FindStringIndex(v)
		structName := v[:ids[0]]
		fieldPath := v[:ids[1]]
		if len(ids) == 0 || (i != 0 && ids[0] != 0) {
			return field, fmt.Errorf("can't get struct name from %v's tag", f.Name)
		}
		var _struct *odmStruct
		// first split is field from col, other field from doc
		if i == 0 {
			_col, err := h.getColByStr(structName)
			if err != nil {
				return field, fmt.Errorf("can't get Col from your register structs")
			}
			_struct = &_col.odmStruct
		} else {
			_struct, err = h.getStructByStr(structName)
			if err != nil {
				return field, errors.New("can't get struct from your register structs")
			}
		}

		f, complexs, err := field.formatField(_struct, fieldPath, i == 0)
		if err != nil {
			return field, fmt.Errorf("can't get struct field from struct: %v", structName)
		}
		goDepend = append(goDepend, f.dependLst...)
		odmDepend = append(odmDepend, f.extendDependLst...)
		complexValues = append(complexValues, complexs...)
	}
	field.goDepend = goDepend
	field.odmDepend = odmDepend
	field.complexValues = complexValues
	return field, nil
}

func (r *HandleField) formatField(o *odmStruct, s string, isFirst bool) (field *structField, complexValues []complexValue, err error) {
	if s == "" {
		err = fmt.Errorf("can't get field use ''")
		return
	}
	sign := s[:1]
	signV := s[1:]
	// two kind of string to get field
	// odmstruct@mark
	// odmstruct.path
	switch sign {
	case "@":
		field = o.getFieldByMark(signV)
		if field == nil || field.complexParent != nil {
			err = fmt.Errorf("can't get field by tag use %v", signV)
			return
		}
	default:
		// mix first split path and orthers path
		// if is not first split, path string wthiout odmstruct name should use "." at first
		if !isFirst && sign != "." {
			err = fmt.Errorf("Can't find field use string %v in struct %v", signV, o.name)
			return
		}
		var path string
		if isFirst {
			path = s
		} else {
			path = signV
		}

		pathLst := strings.Split(path, ".")
		// judge get field by name or by path
		if len(pathLst) == 1 {
			fields := o.getFieldByName(pathLst[0])
			switch len(fields) {
			case 0:
				err = fmt.Errorf("Can't find field use string %v in struct %v", signV, o.name)
				return
			case 1:
				field = fields[0]
				if field == nil || field.complexParent != nil {
					err = fmt.Errorf("can't get field by tag use %v", signV)
					return
				}
			default:
				err = fmt.Errorf("Get to many field use string %v in struct %v", signV, o.name)
				return
			}
		} else {
			fieldLst := o.rootFields
			var complexParent *structField
			lenPath := len(pathLst)
			for i, v := range pathLst {
				// if parent is map or slice ,this value is key of parent kind
				if complexParent != nil {
					complexV := complexValue{
						structId: o.id,
						fieldId:  complexParent.id,
						value:    v,
					}
					complexValues = append(complexValues, complexV)
					complexParent = nil
					continue
				}
				// get field by name
				var checkItem *structField
				for _, fieldLstItem := range fieldLst {
					if fieldLstItem.name == v {
						checkItem = fieldLstItem
						break
					}
				}
				if checkItem == nil {
					err = fmt.Errorf("Can't get field use string '%v' in path %v in struct %v", v, s, o.name)
					return
				}
				if checkItem.isGroupType() {
					complexParent = checkItem
				}
				fieldLst = checkItem.extendChildLst
				if i == lenPath-1 {
					field = checkItem
				}
			}
		}
	}
	return
}

// handleField List
type HandleFieldLst []*HandleField
