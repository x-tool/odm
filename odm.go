package odm

type ODM struct {
	Col    *Col
	DB     *Database // col has db,but it can't use col when handle needless col. Ex: getColLst()
	Handle *handle
	Query  *query
	Result *result
	Err    error
}

func newODM(i interface{}, db *Database, c *Col) *ODM {
	d := &ODM{
		Col:    c,
		DB:     db,
		Handle: newHandle(),
		Query:  newQuery(),
		Result: newResult(i, c),
		Err:    nil,
	}
	return d
}

func (d *ODM) dbName() string {
	return d.DB.name
}
func (d *ODM) colName() string {
	return d.Col.name
}
func (d *ODM) insert() (r interface{}, err error) {
	r, err = d.DB.Dialect.Insert(d)
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
