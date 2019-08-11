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
	runtimeFunctionCalls 
	alias map[string]*odmStruct // register struct alias,not col alias
	handleType
	handleCols // colLst with alias name
	aimer
	writter
	Reader
	Err error
}

func newHandle(db *Database, runtimeFunctionCall *functionCall) *Handle {
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

func (h handleCols) isSingleCol() bool {
	return len(h) == 1
}
