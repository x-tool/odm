package core

type handleType int

// select item from collection, collection like documents and temp documents
type collection struct {
	alias      string
	rootValues []interface{}
}
type collectionLst []*collection

func (cLst collectionLst) getColByAlias(s string) (c []interface{}) {
	for _, v := range cLst {
		if v.alias == s {
			c = v.rootValues
		}
	}
	return
}

func (cLst collectionLst) isSingle() bool {
	return len(cLst) == 1
}

const (
	InsertData handleType = iota
	UpdateData
	DeleteData
	QueryData
)

// handle struct is hock for plugin
type Handle struct {
	handleType
	collectionLst
	aimer
	writter
	reader
	Err error
}

func newHandle() *Handle {
	d := &Handle{}
	return d
}
