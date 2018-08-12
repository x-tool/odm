package core

import (
	"errors"
	"fmt"

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

func (d *Database) getStructByName(name string) (o *odmStruct, err error) {
	o = d.mapStructs[name]
	if o == nil {
		err = errors.New(fmt.Sprintf("Can't find struct name %d in database", name))
	}
	return
}

func (d *Database) RegisterStruct(c interface{}) {
	var _struct *odmStruct
	if v, ok := c.(odmStruct); ok {
		_struct = &v
	} else {
		_struct = newOdmStruct(c)
	}

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
