package core

import (
	"errors"

	"github.com/x-tool/tool"
)

func (d *Database) SyncCols() {
	var colLst []*Col
	for _, v := range d.zoneMap {
		for _, col := range v.ColLst {
			colLst = append(colLst, col)
		}
	}
	d.dialect.SyncCols(colLst)
}

func (d *Database) setHistory() {
	var err error
	d.history = new(history)
	d.history.colNames, err = d.dialect.GetColNames()
	if err != nil {
		tool.Panic("DB", errors.New("Get colNames ERROR"))
	}
}
