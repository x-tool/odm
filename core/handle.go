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
	db *Database
	runtimeFunctionCalls
	alias map[string]*odmStruct // register struct alias,not col alias
	handleType
	usedCols []*odmStruct // used colLst
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
