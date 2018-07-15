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
	aimer
	writter
	reader
	Err error
}

func newHandle(c *Col) *Handle {
	var lst ColLst
	d := &Handle{
		ColLst: append(lst, c),
	}
	return d

}
