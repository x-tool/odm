package core

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/x-tool/tool"
)

func (d *Database) SyncCols() {
	var colLst []*Col
	for _, col := range d.ColLst {
		colLst = append(colLst, col)
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

func (d *Database) GetCol(i interface{}) *Col {
	var name string
	if v, ok := i.(string); !ok {
		name = string(v)
	} else {
		name = reflect.TypeOf(i).Name()
	}
	return d.mapCols[name]
}

var rigisterCols sync.WaitGroup
var rigisterStructs sync.WaitGroup

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.ColLst = append(d.ColLst, _col)
	d.mapCols[_col.Name()] = _col
	d.RegisterStruct(_col.doc.odmStruct)
	rigisterCols.Done()
}

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go d.RegisterCol(v)
	}
	rigisterCols.Wait()
}
