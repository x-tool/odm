package core

type handleType int

const (
	InsertData handleType = iota
	UpdateData
	DeleteData
	QueryData
)

// handle struct is hock for plugin
type Handle struct {
	db    *Database
	alias map[string]*odmStruct // register struct alias,not col alias
	handleType
	handleCols
	aimer
	writter
	Reader
	Err error
}

func newHandle(db *Database) *Handle {
	d := &Handle{
		db: db,
	}
	return d
}

type handleCol struct {
	sign string // alias || col.name
	col  *Col
}

// if new handleCol with alias, could use alias to get field
func newHandleCol(c *Col, alias ...string) *handleCol {
	var name string
	if len(alias) != 0 {
		name = alias[0]
	} else {
		name = c.Name()
	}
	return &handleCol{
		sign: name,
		col:  c,
	}
}

type handleCols []*handleCol

func (hLst *handleCols) add(h *handleCol) {
	l := *hLst
	l = append(l, h)
}
