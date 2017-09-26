package xodm

import "reflect"

type Doc struct {
	Col    *Col
	DB     *Database
	raw    interface{}
	Handle *docHandle
	Err    error
}
type docHandle struct {
	WantFields []*docRootField
	dbName     string
	colName    string
	limit      string
	where      string
}

func newDoc(c *Col) *Doc {
	q := new(docHandle)
	q.colName = c.Name
	q.dbName = c.DB.name
	d := &Doc{
		Col:    c,
		DB:     c.DB,
		raw:    nil,
		Handle: q,
	}
	return d
}

type docRootField struct {
	name       string
	typeName   string
	DBtypeName string
	value      reflect.Value
}

func (d *Doc) insert(i interface{}) (r interface{}, err error) {
	r, err = d.DB.Dialect.Insert(d)
	return
}

func (d *Doc) update(i interface{}) {

}
func (d *Doc) delete(i interface{}) {

}

func (d *Doc) query(i interface{}) {

}

func (d *Doc) Where(s string) *Doc {
	d.Handle.where = s
	return d
}

func (d *Doc) Limit(s string) *Doc {
	d.Handle.limit = s
	return d
}

func (d *Doc) formatQuery() {

}

func (d *Doc) getRootfields() []*docRootField {
	var r []*docRootField
	ivalue := reflect.ValueOf(d.raw)
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
	return r
}
