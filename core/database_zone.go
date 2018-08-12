package core

import (
	"reflect"
	"sync"
)

type zone struct {
	name string
	db   *Database
	ColLst

	mapCols map[string]*Col // use map to get col by name

	aliasFunc func(string) string
}

type zoneLst []*zone

func newZone(name string, d *Database) *zone {
	z := &zone{
		db: d,
	}
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

var rigisterCols sync.WaitGroup
var rigisterStructs sync.WaitGroup

func (z *zone) RegisterCol(c interface{}) {
	_col := newCol(z, c)
	z.ColLst = append(z.ColLst, _col)
	z.mapCols[_col.Name()] = _col
	z.db.RegisterStruct(_col.doc.odmStruct)
	rigisterCols.Done()
}

func (z *zone) RegisterCols(c ...interface{}) {
	for _, v := range c {
		rigisterCols.Add(1)
		go z.RegisterCol(v)
	}
	rigisterCols.Wait()
}
