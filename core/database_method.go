package core

import (
	"errors"
	"sync"

	"github.com/x-tool/tool"
)

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	if _, ok := d.mapCols[_col.Name()]; !ok {
		d.mapCols[_col.Name()] = _col
	}
	d.ColLst = append(d.ColLst, _col)
	rigisterCols.Done()
}

var rigisterCols sync.WaitGroup

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go d.RegisterCol(v)
	}
	rigisterCols.Wait()
}

var rigisterStructs sync.WaitGroup

func (d *Database) RegisterStruct(c interface{}) {
	_struct := newOdmStruct(c)
	if _, ok := d.mapStructs[_struct.allName]; !ok {
		d.mapStructs[_struct.allName] = _struct
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

func (d *Database) getColByName(name string) *Col {
	return d.mapCols[name]
}
