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
func (d *Database) NewCol(cols ...ColInterface) error {
	var err error
	for _, col := range cols {
		err = d.newdatabaseCol(col)
	}
	return err
}
