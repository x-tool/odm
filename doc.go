package odm

type Doc struct {
	Col    *Col
	DB     *Database // col has db,but it can't use col when handle needless col. Ex: getColLst()
	Handle *handle
	Query  *query
	Result *result
	Err    error
}

func newDoc(i interface{}, db *Database, c *Col) *Doc {
	d := &Doc{
		Col:    c,
		DB:     db,
		Handle: newHandle(),
		Query:  newQuery(),
		Result: newResult(i, c),
		Err:    nil,
	}
	return d
}

func (d *Doc) dbName() string {
	return d.DB.name
}
func (d *Doc) colName() string {
	return d.Col.name
}
func (d *Doc) insert() (r interface{}, err error) {
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
	// d.Handle.where = s
	return d
}

func (d *Doc) Limit(s string) *Doc {
	// d.Handle.limit = s
	return d
}
