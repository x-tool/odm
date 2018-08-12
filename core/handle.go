package core

type handleType int

type handleCol struct {
	sign string // alias || col.name
	col  *Col
}
type handleCols []*handleCol

const (
	InsertData handleType = iota
	UpdateData
	DeleteData
	QueryData
)

// handle struct is hock for plugin
type Handle struct {
	db *Database
	handleType
	handleCols
	aimer
	writter
	reader
	Err error
}

func newHandle() *Handle {
	d := &Handle{}
	return d
}
