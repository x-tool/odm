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
	DB        *Database
	Name      string
	detailLst []*ColDetail
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

func (c *Col) getRootDetails() (colLst []*ColDetail) {
	for _, v := range c.detailLst {
		if v.Pid == -1 {
			colLst = append(colLst, v)
		}
	}
	return
}

type ColDetail struct {
	Name   string
	Type   string
	DBType string
	Id     int
	Pid    int
}

func NewCol(db *Database, i ColInterface) *Col {
	c := new(Col)
	c.Name = i.ColName()
	c.DB = db
	c.setDetail(i)
	db.ColLst = append(db.ColLst, c)
	return c
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
		if FieldType.Kind() == reflect.Struct {
			mergeDetail(c, FieldType, id)
		} else {
			c.detailLst = append(c.detailLst, &ColDetail{
				Name:   FieldName,
				Type:   FieldType.Kind().String(),
				DBType: FieldDBType,
				Id:     id,
				Pid:    Pid,
			})
		}

	}
}
