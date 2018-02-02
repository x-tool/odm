package odm

import "reflect"

type ColInterface interface {
	ColName() string
}

type Col struct {
	name           string
	hasDocModel    bool
	DocModel       string
	hasDeleteField bool
	deleteField    string
	Doc            *Doc
}

func NewCol(i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.DB = d
	c.Doc = NewDoc(c, i)
	c.hasDocModel, c.DocModel = c.Doc.DocModel()
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
