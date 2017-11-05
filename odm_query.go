package odm

import (
	"errors"
	"reflect"
	"strings"

	"github.com/x-tool/tool"
)

type query struct {
	Col       *Col
	dependLst []*DocField
	queryKind string
	queryV    *reflect.Value
	modeV     *reflect.Value
	queryLst  []queryItem
	limitNum  int
	limitDesc bool
	querySet  []queryItem
}

type queryItem struct {
	queryRootField
	whereCheck string
	whereAnd   bool
}

func newQuery(rV *reflect.Value, c *Col, t string) *query {
	r := &query{
		Col:       c,
		queryV:    rV,
		queryKind: t,
	}
	r.setDependToDoc()
	return r
}
func newqueryWithoutCol(rV *reflect.Value) *query {
	r := &query{
		queryV: rV,
	}
	return r
}

func (r *query) setDependToDoc() {
	T := r.queryV.Type()
	var value reflect.Type
	if T.Kind() == reflect.Slice {
		value = T.Elem()
	} else {
		value = T
	}
	var valueItem reflect.Value
	var valueItemT reflect.Type
	if value.Kind() == reflect.Slice {
		valueItem = r.queryV.Elem()
	} else {
		valueItem = *r.queryV
	}
	valueItemT = valueItem.Type()
	for i := 0; i < valueItem.NumField(); i++ {
		field := valueItem.Field(i)
		fieldT := valueItemT.Field(i)
		if isDocMode(fieldT.Name) {
			r.modeV = &field
		}
		newqueryItem := r.DependToDoc(strings.Split(fieldT.Tag.Get(tagName), "."), fieldT.Name)
		r.dependLst = append(r.dependLst, newqueryItem)
	}

}
func (r *query) dependtoDocOneStr(s string) (d *DocField) {
	_s := strings.Split(s, ".")
	if len(_s) > 1 {
		return r.DependToDoc(_s[:len(_s)-2], _s[len(_s)-1])
	} else {
		return r.DependToDoc([]string{}, _s[0])
	}
}
func (r *query) DependToDoc(dependLst []string, name string) (d *DocField) {
	if len(dependLst) == 0 {
		field := r.Col.Doc.getFieldByName(name)
		if len(field) != 1 {
			tool.Panic("ODM", errors.New("name not be single, you should add dependLst to find doc field"))
		} else {
			return field[0]
		}
	} else {
		docFieldLst := r.Col.Doc.getFieldByName(name)
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
