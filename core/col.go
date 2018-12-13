package core

type Col struct {
	db    *Database
	name  string
	alias string // name to database
	mode  *StructField
	odmStruct
}

func (c *Col) Name() string {
	return c.name
}

func newCol(d *Database, i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.db = d
	c.odmStruct = *newOdmStruct(i)
	c.mode = c.findDocModeField()
	return c
}
