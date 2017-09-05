package xodm

const (
	tagName = "xodm"
)

type Database struct {
	ColLst []*Col
	Dialect
}

func (d *Database) NewConn() (c Conn, err error) {
	return d.Conn()
}
func (d *Database) SyncCols(cols ...ColInterface) {
	for _, col := range cols {
		d.syncCol(col)
	}
}

func (d *Database) Insert(c ColInterface) {

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
