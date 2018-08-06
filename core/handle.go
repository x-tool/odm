package core

type handleType int
type collection interface{}

const (
	InsertData handleType = iota
	UpdateData
	DeleteData
	QueryData
)

// handle struct is hock for plugin
type Handle struct {
	handleType
	collectionNamesMap map[string]collection
	aimer
	writter
	reader
	Err error
}

func newHandle() *Handle {
	d := &Handle{}
	return d
}
