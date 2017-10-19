package odm

import "reflect"

type ODM struct {
	Col    *Col
	DB     *Database // col has db,but it can't use col when handle needless col. Ex: getColLst()
	Handle *handle
	Query  *query
	Result *result
	R      *reflect.Value
	Err    error
}

func newODM(db *Database, c *Col) *ODM {
	d := &ODM{
		Col: c,
		DB:  db,
	}
	return d
}

func newODMwithoutCol(db *Database) *ODM {
	d := &ODM{
		DB: db,
	}
	return d
}

func (d *ODM) dbName() string {
	return d.DB.name
}
func (d *ODM) colName() string {
	return d.Col.name
}
func (d *ODM) insert(i interface{}) (err error) {
	r := reflect.ValueOf(i)
	d.R = &r
	d.Handle = newHandle(HandleInsert)
	d.Query = newQuery(&r, d.Col)
	d.Result = newResult(&r, d.Col)
	err = d.DB.Dialect.Insert(d)
	return
}

func (d *ODM) update(i interface{}) {

}
func (d *ODM) delete(i interface{}) {

}

func (d *ODM) query(i interface{}) {

}

func (d *ODM) Where(s string) *ODM {
	// d.Handle.where = s
	return d
}

func (d *ODM) Limit(s string) *ODM {
	// d.Handle.limit = s
	return d
}

func (d *ODM) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}
