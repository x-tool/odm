package core

import "reflect"

type HandleType int

const (
	InsertData HandleType = iota
	addData
	updateData
	deleteData
)

type Handle struct {
	col         *Col
	target      *docField
	Query       *query
	Result      *result
	OriginValue *reflect.Value
	OriginType  reflect.Type
	Err         error
}

func newHandle(c *Col) *Handle {
	d := &Handle{
		col: c,
	}
	return d

}

func (d *Handle) dbName() string {
	return d.target.doc.col.database.name
}

func (d *Handle) colName() string {
	return d.target.doc.col.name
}

func (d *Handle) insert(i interface{}) (err error) {
	r := reflect.Indirect(reflect.ValueOf(i))
	d.OriginValue = &r

	d.Query = newQuery(&r, d, "insert")
	d.Result = newResult(&r, d.col)
	// modeInsert(d)
	err = d.col.database.dialect.Insert(d) //.handle(d)
	return
}

func (d *Handle) update(i interface{}) {

}

func (d *Handle) delete(err error) {
	// if d.Col.doc.getDeleteFieldName() != "" {
	// 	err = d.DB.Dialect.Update(d)
	// } else {
	// 	err = d.DB.Dialect.Delete(d)
	// }

	// return
}

func (d *Handle) get(i interface{}) {

}

func (d *Handle) Where(s string) *Handle {
	// d.Handle.where = s
	return d
}

func (d *Handle) Limit(s string) *Handle {
	// d.Handle.limit = s
	return d
}

func (d *Handle) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}
