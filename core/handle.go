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

func (h HandleFilter) Kind() Kind {
	return h.target.GetKind()
}
func (h HandleFilter) FieldName() string {
	return h.target.GetName()
}
func (h HandleFilter) Vakue() interface{} {
	return h.value
}

type HandleFilterLst []*HandleFilter

type HandleSetValue struct {
	target     *docField
	handleType handleType
	value      interface{}
}

func (h HandleSetValue) Kind() Kind {
	return h.target.GetKind()
}
func (h HandleSetValue) FieldName() string {
	return h.target.GetName()
}
func (h HandleSetValue) Vakue() interface{} {
	return h.value
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
	Origin
	Err error
}

func (d *Handle) GetDBName() string {
	return d.col.database.name
}

func (d *Handle) GetColName() string {
	return d.col.name
}

func (d *Handle) insert(db *Database, i interface{}) (err error) {
	d.setOrigin(i)
	d.col = db.GetColByName(GetColNameByReflectType(d.GetOrigin().Type()))

	err = db.dialect.Insert(d)
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

func newHandle(c *Col) *Handle {
	d := &Handle{
		col: c,
	}
	return d

}

type Origin struct {
	result
	OriginValue *reflect.Value
	OriginType  reflect.Type
}

func (o *Origin) setOrigin(i interface{}) {
	value := reflect.Indirect(reflect.ValueOf(i))
	o.OriginValue = &value
	o.OriginType = o.OriginValue.Type()
}

func (o *Origin) GetOrigin() *Origin {
	return o
}

func (o *Origin) Value() *reflect.Value {
	return o.OriginValue
}

func (o *Origin) Type() reflect.Type {
	return o.OriginType
}
