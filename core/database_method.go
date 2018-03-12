package core

import (
	"errors"
	"sync"

	"github.com/x-tool/tool"
)

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	if _, ok := d.mapCols[_col.GetName()]; !ok {
		d.mapCols[_col.GetName()] = _col
	}
	d.ColLst = append(d.ColLst, _col)
	syncLock.Done()
}

var syncLock sync.WaitGroup

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		syncLock.Add(1)
		go d.RegisterCol(v)
	}
	syncLock.Wait()
}

func (d *Database) SyncCols() {
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

func (d *Database) getColByName(name string) *Col {
	return d.mapCols[name]
}
