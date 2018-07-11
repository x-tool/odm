package core

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type zone struct {
	name       string
	mapCols    map[string]*Col       // use map to get col by name
	mapStructs map[string]*odmStruct // use map to get structs by name, I think struct name should be unique where ever package, if not user should write whole pkgPath and name in one string to get one struct
	nameFunc   func(string) string
}

type zoneLst []*zone

func newZone(name string) {

}

func (z *zone) GetCol(i interface{}) *Col {
	var name string
	if v, ok := i.(string); !ok {
		name = string(v)
	} else {
		name = reflect.TypeOf(i).Name()
	}
	return z.mapCols[name]
}

func (z *zone) getStructByName(name string) (o *odmStruct, err error) {
	o = z.mapStructs[name]
	if o == nil {
		err = errors.New(fmt.Sprintf("Can't find struct name %d in database", name))
	}
	return
}

var rigisterCols sync.WaitGroup
var rigisterStructs sync.WaitGroup

func (z *zone) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.ColLst = append(d.ColLst, _col)
	d.mapCols[_col.Name()] = _col
	d.RegisterStruct(_col.doc.odmStruct)
	rigisterCols.Done()
}

func (z *zone) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go d.RegisterCol(v)
	}
	rigisterCols.Wait()
}

func (z *zone) RegisterStruct(c interface{}) {
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

func (z *zone) RegisterStructs(c ...interface{}) {
	for _, v := range c {
		rigisterStructs.Add(1)
		go d.RegisterStruct(v)
	}
	rigisterStructs.Wait()
}
