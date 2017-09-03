package xodm

import (
	"reflect"
)

type ColLst []Col

type ColInterface interface {
	Name() string
}

type Col struct {
	Name      string
	detailLst []*ColDetail
}

func (c *Col)mergeDetail(v reflect.Value, Id){
	p := reflect.ValueOf(col)
	v := p.Elem()
	t := v.Type()
	if v.Kind() == reflect.Struct {
		col := new(Col)
		col.Name = t.Name()
		colFieldNum := v.NumField()
		// make ColLst in a col
		for i := 0; i < colFieldNum; i++ {
			field := t.Field(i)
			FieldName := field.Name
			FieldTag := field.Tag.Get(tagName)
			FieldType := field.Type()
			FieldDBType := d.SwitchType(FieldTag)
			if FieldTag == "" {
				continue
			}
			col.detailLst = append(col.detailLst, &ColDetail{
				Name: FieldName,
				Type: FieldType,
				DBtype: FieldDBType,
				Id: 
			})
		}
		d.mergeCol(colName, colFieldLst)
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+v.Kind().String()+"!,Type should be Struct"))
	}
	return nil
}

type ColDetail struct {
	Name   string
	Type   string
	DBtype string
	Id     int
	Pid    int
}

func NewCol(i ColInterface)( *Col){
	c := new(Col)
	c.mergeDetail(i)
	return c
}
