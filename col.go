package odm

import "github.com/x-tool/odm/module/model"

type col struct {
	baseCol *model.Col
}

type colLst []*col

func newCol(d *database, i interface{}) *col {
	_col := new(col)
	_col.baseCol = model.NewCol(i)
}
