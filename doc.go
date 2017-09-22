package xodm

import "reflect"

type Question struct {
	dbName  string
	colName string
	limit   string
	where   string
}

type Doc struct {
	Col      *Col
	DB       *Database
	raw      interface{}
	Answer   interface{}
	Question *Question
	Err      error
}

func newDoc(c *Col) *Doc {
	d := &Doc{
		Col:      c,
		DB:       c.DB,
		raw:      nil,
		Answer:   nil,
		Question: new(Question),
	}
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

func (d *Doc) formatRaw() {
	t := reflect.TypeOf(d.raw)
	newV := reflect.New(t)

}

func (d *Doc) insert(i interface{}) {

}

func (d *Doc) update(i interface{}) {

}
func (d *Doc) delete(i interface{}) {

}

func (d *Doc) query(i interface{}) {

}
