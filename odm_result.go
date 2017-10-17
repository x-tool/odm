package odm

import (
	"reflect"
	"time"

	"github.com/x-tool/tool"
)

type result struct {
	Doc       *Doc
	resultLst []*resultItem
	raw       interface{}
}
type resultItem struct {
}

func newResult(i interface{}, c *Col) *result {
	r := &result{
		Doc: c.Doc,
		raw: i,
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
		tool.Panic("ODM", "have not Ptr, Can't write In")
	} else {
		value = T.Elem()
	}
	if value.Kind() == reflect.Slice {

	}
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
