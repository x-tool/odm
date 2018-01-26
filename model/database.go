package model

import "github.com/x-tool/odm/module/model"

// database use
type database struct {
	name string
	colLst
}

type colLst []*model.Col

func NewDatabase(name string) *database {
	_d := new(database)
	_d.Name = name
	return _d
}

func (d *database) GetName() string {
	return d.name
}

func (d *database) RegisterCol(c interface{}) {
	_col := NewCol(d, c)
	d.colLst = append(d.colLst, _col)
}

func (d *database) RegisterCols(c ...interface{}) {
	for i := range c {
		go d.RegisterCol(i)
	}
}

func (d *database) GetColByName(name string) *Col {
	var col *Col
	for _, v := range d.colLst {
		if v.name == name {
			col = v
			break
		}
	}
	return col
}

func (d *database) GetName() string {
	return d.name
}
