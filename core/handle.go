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
	col *Col
	aimer
	writter
	reader
	Err error
}

func newHandle(c *Col) *Handle {
	d := &Handle{
		db:      c.database,
		col:     c,
		context: con,
	}
	return d

}
