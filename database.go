package xodm

const (
	tagName = "xodm"
)

type Database struct {
	name   string
	ColLst []*Col
	Dialect
}

func (d *Database) NewConn() (c Conn, err error) {
	return d.Conn()
}
func (d *Database) SyncCols(cols ...interface{}) {
	for _, col := range cols {
		d.syncCol(col)
	}
}

func (d *Database) Insert(r interface{}, doc *Doc) {
	d.Dialect.Insert(r, doc)
}
func (d *Database) Update(c interface{}) {

}
func (d *Database) Delete(c interface{}) {

}
func (d *Database) Query(c interface{}) {

}

func (d *Database) getCol(name string) *Col {
	var col *Col
	for _, v := range d.ColLst {
		if v.Name == name {
			col = v
			break
		}
	}
	return col
}
