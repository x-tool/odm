package odm

import "github.com/x-tool/odm/model"

type database struct {
	baseDB model.Database
	colLst
}

func newDatabase(s string) *database {
	_d := new(database)
	_d.baseDB = model.NewDatabase(s)
	return _d
}

func (d *database) RegisterCol(i interface{}) {
	col := newCol(d, i)
	d.colLst = append(d.colLst, col)
}

func (d *database) RegisterCols(i ...Interface{}) {
	for _,v:=range i{
		go d.RegisterCol(v)
	}
}

func (d *database) GetDBName()string{
	return d.baseDB.getName()
}