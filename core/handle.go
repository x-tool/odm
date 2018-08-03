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
	handleType
	ColLst
	mapColsNameLst map[string]*Col // allow aliasname
	aimer
	writter
	reader
	Err error
}

func newHandle() *Handle {
	d := &Handle{}
	return d
}
