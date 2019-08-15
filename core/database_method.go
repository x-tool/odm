package core

import (
	"fmt"
	"reflect"
	"sync"
)

var rigisterCols sync.WaitGroup
var rigisterStructs sync.WaitGroup

func (d *Database) getStructByName(name string) (o *odmStruct, err error) {
	o = d.mapStructs[name]
	if o == nil {
		err = fmt.Errorf("Can't find struct name %v in database", name)
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
	rigisterStructs.Done()
}

func (d *Database) RegisterStructs(c ...interface{}) {
	for _, v := range c {
		rigisterStructs.Add(1)
		go d.RegisterStruct(v)
	}
	rigisterStructs.Wait()
}

func (d *Database) GetCol(i interface{}) (c *Col, err error) {
	var name string

	if v, ok := i.(string); ok {
		name = string(v)
	} else {
		name = reflect.TypeOf(i).Name()
	}
	c = d.mapCols[name]

	if c == nil {
		err = fmt.Errorf("can't get Col By Name %v", name)
	}
	return
}

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.ColLst = append(d.ColLst, _col)
	d.mapCols[_col.Name()] = _col
	rigisterStructs.Add(1)
	d.RegisterStruct(_col.odmStruct)
	rigisterCols.Done()
}

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go d.RegisterCol(v)
	}
	rigisterCols.Wait()
}

func (d *Database) SyncCols() {
	d.dialect.SyncCols(d.ColLst)
}

// rigister user custom data type
func (d *Database) RegisterType(name string, value interface{}, funcs customTypeInterface) {
	c := newCustomType(name, value, funcs)
	customTypeBox.typeLst = append(customTypeBox.typeLst, c)
}

func (d *Database) newHandle() (h *Handle) {
	return newHandle(d)
}
