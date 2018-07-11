package core

import (
	"errors"

	"github.com/x-tool/tool"
)

func (d *Database) SyncCols() {
	for _, v := range d.ColLst {
		col := *v
		col.alias = d.config.colNameAlias(col.name)
	}
	d.dialect.SyncCols(d.ColLst)
}

func (d *Database) setHistory() {
	var err error
	d.history = new(history)
	d.history.colNames, err = d.dialect.GetColNames()
	if err != nil {
		tool.Panic("DB", errors.New("Get colNames ERROR"))
	}
}

func (d *Database) ColNameAlias(f func(string) string) {
	if d.isSyncCols {
		tool.Panic("DB", errors.New("It can't works, because database has been sync, you should write ColNameAlias method before SyncCols method"))
	}
	d.config.colNameAlias = f
}
