package core

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type HandleField struct {
	handle        *Handle
	col           *Col
	depend        dependLst
	complexValues []string //  slice id or map key,if child field is not complex type add nil
}

func (c HandleField) getValue(id int) string {
	return c.complexValues[id]
}

// if handle by only one col
// path|structname.path|......
// else
// colname.path|......
func newHandleField(h *Handle, s string) (field *HandleField, err error) {
	field = &HandleField{
		handle: h,
	}
	var goDepend dependLst
	var complexValues []string
	// format StructFieldTag
	odmLst := strings.Split(s, "|")
	splitStructReg := regexp.MustCompile(`\w`)
	for i, v := range odmLst {
		var structName string
		var fieldPath string
		var f *StructField
		var complexs []string
		// first path for col
		// other for struct
		if i == 0 {
			var col *Col
			if h.isSingleCol() {
				col = h.GetCol()
				field.col = col
				structName = col.name
				fieldPath = v
			} else {
				regId := splitStructReg.FindStringIndex(v)
				structName = v[:regId[0]]
				fieldPath = v[regId[0]:]
				col, err = h.getColByStr(structName)
				field.col = col
				if len(regId) == 0 || err != nil {
					return field, fmt.Errorf("can't get col name from %v's tag", f.Name())
				}

			}
			f, complexs, err = field.formatField(&col.odmStruct, fieldPath)
		} else {
			f, complexs, err = field.formatFieldWithStructName(v)
		}

		if err != nil {
			return field, fmt.Errorf("can't get struct field from struct: %v", structName)
		}
		goDepend = append(goDepend, f.dependLst...)
		complexValues = append(complexValues, complexs...)
	}
	field.depend = goDepend
	field.complexValues = complexValues
	return field, nil
}

func (h *HandleField) formatFieldWithStructName(s string) (field *StructField, complexValues []string, err error) {
	var _struct *odmStruct
	reg := regexp.MustCompile(`\w`)
	ids := reg.FindStringIndex(s)
	structName := s[:ids[0]]
	fieldPath := s[:ids[1]]
	_struct, err = h.handle.getStructByStr(structName)
	if err != nil {
		return nil, nil, errors.New("can't get struct from your register structs")
	}

	return h.formatField(_struct, fieldPath)
}

// format field from path
func (r *HandleField) formatField(o *odmStruct, s string) (field *StructField, complexValues []string, err error) {
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
		if sign != "." {
			err = fmt.Errorf("Can't find field use string %v in struct %v", signV, o.name)
			return
		}
		var path string
		path = signV

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
			var complexParent *StructField
			lenPath := len(pathLst)
			for i, v := range pathLst {
				// if parent is map or slice ,this value is key of parent kind
				if complexParent != nil {
					complexValues[i] = v
					complexParent = nil
					continue
				}
				// get field by name
				var checkItem *StructField
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
