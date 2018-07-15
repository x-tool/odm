package core

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type zone struct {
	name string
	ColLst
	odmStructLst
	mapCols    map[string]*Col       // use map to get col by name
	mapStructs map[string]*odmStruct // use map to get structs by name, I think struct name should be unique where ever package, if not user should write whole pkgPath and name in one string to get one struct
	aliasFunc  func(string) string
}

type zoneLst []*zone

func newZone(name string) *zone {
	z := &zone{}
	// col.alias = v.aliasFunc(col.name)
	return z
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
	_col := newCol(z, c)
	z.ColLst = append(z.ColLst, _col)
	z.mapCols[_col.Name()] = _col
	z.RegisterStruct(_col.doc.odmStruct)
	rigisterCols.Done()
}

func (z *zone) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go z.RegisterCol(v)
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

	if _, ok := z.mapStructs[_struct.name]; !ok {
		z.mapStructs[_struct.name] = _struct
	}
	z.odmStructLst = append(z.odmStructLst, _struct)
	rigisterCols.Done()
}

func (z *zone) RegisterStructs(c ...interface{}) {
	for _, v := range c {
		rigisterStructs.Add(1)
		go z.RegisterStruct(v)
	}
	rigisterStructs.Wait()
}
