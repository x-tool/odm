package odm

import (
	"reflect"
)

type ColInterface interface {
	ColName() string
}

type Col struct {
	name           string
	hasDeleteField bool
	deleteField    string
	Doc
	colModeJ model.ColModer
}

func newCol(i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.DB = d
	c.Doc = NewDoc(c, i)
	c.deleteField = c.Doc.getDeleteFieldName()
	if c.deleteField != "" {
		c.hasDeleteField = true
	}
	return c
}

func GetColName(i interface{}) (name string) {
	if colI, ok := i.(ColInterface); ok {
		name = colI.ColName()
	} else {
		v := reflect.TypeOf(i)
		if v.Kind() == reflect.Ptr {
			name = v.Elem().Name()
		} else {
			name = v.Name()
		}

	}
	return
}
