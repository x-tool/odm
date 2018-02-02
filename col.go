package odm

import (
	"reflect"
)

type ColInterface interface {
	ColName() string
}

type Col struct {
	database *database
	name     string
	Doc
	colModeJ model.ColModer
}

func newCol(db *database, i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.db = db
	c.Doc = NewDoc(c, i)
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
