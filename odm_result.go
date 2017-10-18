package odm

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/x-tool/tool"
)

type result struct {
	Doc       *Doc
	resultLst []*DocField
	odm       *ODM
}

func newResult(rV *reflect.Value, c *Col) *result {
	r := &result{
		Doc: c.Doc,
		odm: i,
	}
	r.format()
	return r
}
func newResultWithoutCol(i interface{}) *result {
	r := &result{
		Doc: nil,
		raw: i,
	}
	return r
}

func (r *result) format() {
	T := reflect.TypeOf(r.raw)
	var value reflect.Type
	if T.Kind() != reflect.Ptr {
		tool.Panic("ODM", errors.New("have not Ptr, Can't write In"))
	} else {
		value = T.Elem()
	}
	if value.Kind() == reflect.Slice {
		valueItem := value.Elem()
		for i := 0; i < valueItem.NumField(); i++ {
			field := valueItem.Field(i)
			newResultItem := r.DependToDoc(field.Tag.Get(tagName), field.Name)
			r.resultLst = append(r.resultLst, newResultItem)
		}
	}
}

func (r *result) DependToDoc(tag string, name string) (d *DocField) {
	if tag == "" {
		field := r.Doc.getFieldByName(name)
		if len(field) != 1 {
			tool.Panic("ODM", errors.New("name not be single, you should add tag to find doc field"))
		} else {
			return field[0]
		}
	} else {
		dependLst := strings.Split(tag, ".")
		docFieldLst := r.Doc.getFieldByName(name)
		for _, val := range docFieldLst {
			if len(dependLst) != len(val.dependLst) {
				continue
			}
			var check bool = true
			for i, _ := range val.dependLst {
				if val.dependLst[i].Name != dependLst[i] {
					check = false
					break
				}
			}
			if check {
				d = val
				break
			}
		}
	}
	return
}
func (r *result) selectValidFields(dLst []*docRootField) (vLst []*docRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}

func (r *result) checkZero(v reflect.Value) bool {

	var isValid bool
	value := v.Interface()
	switch v.Kind() {
	case reflect.String:
		if value.(string) == "" {
			isValid = true
		}
	case reflect.Bool:
		if value.(bool) {
			isValid = true
		}
	case reflect.Int:
		if value.(int) == 0 {
			isValid = true
		}
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		if v.Len() != 0 {
			isValid = true
		}
	case reflect.Struct:
		if _v, ok := value.(time.Time); ok {
			if !_v.IsZero() {
				isValid = true
			}
		}
	default:
		isValid = false
	}
	return isValid
}
