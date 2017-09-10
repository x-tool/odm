package xodm

import "reflect"

type docfield struct {
	
	Name    string
	Type    string
	DBType  string
	Id      int
	Pid     int
	isRoot  bool
	Tag     string
	funcLst map[string]string
}

type Doc struct {
	DB *Database
	Col *Col
	fields []*docfield
}

func (c *Doc) getRootDetails() (doc []*docfield) {
	for _, v := range c.fields {
		if v.Pid == -1 {
			doc = append(doc, v)
		}
	}
	return
}

func NewDoc(c *Col, i ColInterface) *Doc {
	doc := new(Doc)
	doc.Col = c
	doc.DB = c.DB
	fieldLst := reflect.TypeOf(i)
	cont := fieldLst.NumField()
	for i := 0; i < cont; i++ {
		field := fieldLst.Field(i)
		NewDocField(doc, &field.Type, -1)
	}

}

func NewDocField(d *Doc, t *reflect.Type, Pid int) {
	field := new(docfield)
	field.Name = t.Name
	field.Type = t.Type.Kind().String()
	field.Tag = t.Tag.Get(tagName)
	field.DBType
	if 
}
