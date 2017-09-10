package xodm

type ColInterface interface {
	ColName() string
}

type Col struct {
	DB   *Database
	Name string
	Doc  *Doc
}

func NewCol(d *Database, i ColInterface) *Col {
	c := new(Col)
	c.Name = i.ColName()
	c.DB = d
	c.Doc = NewDoc(c, i)
	return c
}
