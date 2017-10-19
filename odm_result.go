package odm

import (
	"errors"
	"reflect"
	"strings"

	"github.com/x-tool/tool"
)

type result struct {
	Doc            *Doc
	resultFieldLst []*DocField
	resultV        *reflect.Value
}

func newResult(rV *reflect.Value, c *Col) *result {
	r := &result{
		Doc:     c.Doc,
		resultV: rV,
	}
	r.format()
	return r
}
func newResultWithoutCol(rV *reflect.Value) *result {
	r := &result{
		Doc:     nil,
		resultV: rV,
	}
	return r
}

func (r *result) format() {
	T := r.resultV.Type()
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
			r.resultFieldLst = append(r.resultFieldLst, newResultItem)
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
