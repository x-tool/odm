package xodm

import "reflect"

type Question struct{}
type Answer struct{}
type Doc struct {
	Col      *Col
	DB       *Database
	raw      interface{}
	Answer   *Answer
	Question *Question
}

func newDoc(c *Col, i interface{}) *Doc {
	d := &Doc{
		Col:      c,
		DB:       c.DB,
		raw:      i,
		Answer:   new(Answer),
		Question: new(Question),
	}
	d.formatRaw()
	return d
}

type docRootField struct {
	name       string
	typeName   string
	DBtypeName string
	value      interface{}
}

func (d *Doc) getRootfields() (r []*docRootField) {
	ivalue := reflect.ValueOf(d.raw).Elem()
	rootDetails := d.Col.OriginDocs.getRootDetails()
	for _, v := range rootDetails.getRootSinpleFields() {
		f := &docRootField{
			name:       v.Name,
			typeName:   v.Type,
			DBtypeName: v.DBType,
			value:      ivalue.FieldByName(v.Name),
		}
		r = append(r, f)
	}
	for _, v := range rootDetails.getRootComplexFields() {
		fields := d.Col.OriginDocs.getChildFields(v)
		for _, val := range fields {
			f := &docRootField{
				name:       val.Name,
				typeName:   val.Type,
				DBtypeName: val.DBType,
				value:      ivalue.FieldByName(val.Name),
			}
			r = append(r, f)
		}
	}
	return
}
