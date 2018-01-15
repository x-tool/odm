package model

import (
	"github.com/x-tool/tool"
)

type database = model.Database

func newDatabase(name string) *database {
	_d := new(database)
	_d.Name = name
	return _d
}

func (d *database) SyncCols(cols ...interface{}) {
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

func (d *database) GetCol(name string) *Col {
	var col *Col
	for _, v := range d.ColLst {
		if v.name == name {
			col = v
			break
		}
	}
	return col
}
