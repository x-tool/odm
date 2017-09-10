package xodm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type ColInterface interface {
	ColName() string
}

type Col struct {
	DB   *Database
	Name string
	Doc
}

func (c *Col) setDetail(col ColInterface) {
	p := reflect.ValueOf(col)
	v := p.Elem()
	if v.Kind() == reflect.Struct {
		t := v.Type()
		mergeDetail(c, t, -1)
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+v.Kind().String()+"!,Type should be Struct"))
	}
}

func (c *Col) switchType(s string) string {
	return c.DB.SwitchType(s)
}
func mergeDetail(c *Col, t reflect.Type, Pid int) {

	colFieldNum := t.NumField()
	// make ColLst in a col
	for i := 0; i < colFieldNum; i++ {
		field := t.Field(i)
		FieldName := field.Name
		FieldTag := field.Tag.Get(tagName)
		if FieldTag == "" {
			continue
		}
		FieldType := field.Type
		FieldDBType := c.DB.SwitchType(FieldTag)
		id := len(c.detailLst)
		c.detailLst = append(c.detailLst, &ColDetail{
			Name:   FieldName,
			Type:   FieldType.Kind().String(),
			DBType: FieldDBType,
			Id:     id,
			Pid:    Pid,
		})
		// range struct
		if FieldType.Kind() == reflect.Struct {
			mergeDetail(c, FieldType, id)
		}
	}
}
