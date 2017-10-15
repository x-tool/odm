package odm

import (
	"github.com/x-tool/tool"
)

const (
	tagName = "odm"
)

type Database struct {
	Client           *Client
	name             string
	activeColNameLst []string
	ColLst           []*Col
	Dialect
}

// func (d *Database) NewConn() (c Conn, err error) {
// 	return d.Conn()
// }
func (d *Database) SyncCols(cols ...interface{}) {
	activeCols, err := d.GetColNames(d)
	if err != nil {
		tool.Panic("DB", err)
	}
	d.activeColNameLst = activeCols
	for _, col := range cols {
		// colName := GetColName(col)
		// if d.findColINactiveCol(colName) {

		// } else {
		d.syncCol(col)
		// }

	}
}

func (d *Database) getCol(name string) *Col {
	var col *Col
	for _, v := range d.ColLst {
		if v.name == name {
			col = v
			break
		}
	}
	return col
}
