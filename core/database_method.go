package core

import (
	"errors"
	"sync"

	"github.com/x-tool/tool"
)

var rigisterCols sync.WaitGroup
var rigisterStructs sync.WaitGroup

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.ColLst = append(d.ColLst, _col)
	d.mapCols[_col.Name()] = _col
	rigisterCols.Done()
}

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go d.RegisterCol(v)
	}
	rigisterCols.Wait()
}

func (d *Database) RegisterStruct(c interface{}) {
	_struct := newOdmStruct(c)
	if _, ok := d.mapStructs[_struct.name]; !ok {
		d.mapStructs[_struct.name] = _struct
	}
	d.odmStructLst = append(d.odmStructLst, _struct)
	rigisterCols.Done()
}

func (d *Database) RegisterStructs(c ...interface{}) {
	for _, v := range c {
		rigisterStructs.Add(1)
		go d.RegisterStruct(v)
	}
	rigisterStructs.Wait()
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

func (d *Database) ColNameAlias(f func(string) string) {
	d.config.colNameAlias = f
}
