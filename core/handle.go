package core

import "reflect"

type handleType int

const (
	insertData handleType = iota
	updateData
	deleteData
)

type contrast int

const (
	sameData      contrast = iota // like
	sameDataLeft                  // ??like
	sameDataRight                 // like??
	equalData                     // ==
)

type HandleFilter struct {
	target   *docField
	contrast contrast
	value    interface{}
}
type HandleFilterLst []*HandleFilter

type HandleSetValue struct {
	target     *docField
	handleType handleType
	value      interface{}
}
type HandleSetValueLst []*HandleSetValue

type HandleGroup struct {
	filterLst HandleFilterLst
	setLst    HandleSetValue
}

type Handle struct {
	// ptr to col
	col            *Col
	filterCols     HandleFilterLst
	HandleGroupLst []*HandleGroup
	Result         *result
	OriginValue    *reflect.Value
	OriginType     reflect.Type
	Err            error
}

func newHandle(c *Col) *Handle {
	d := &Handle{
		col: c,
	}
	return d

}

func (d *Handle) GetDBName() string {
	return d.col.database.name
}

func (d *Handle) GetColName() string {
	return d.col.name
}

func (d *Handle) insert(db *Database, i interface{}) (err error) {
	r := reflect.Indirect(reflect.ValueOf(i))
	d.col = db.GetColByName(GetColNameByReflectType(r.Type()))
	d.OriginValue = &r

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
