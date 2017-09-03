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
