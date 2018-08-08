package core

type handleType int

// select item from collection, collection like documents and temp documents
type collection struct {
	alias string
	structFieldLst
}
type collectionLst []*collection

func (cLst collectionLst) getColByAlias(s string) (c structFieldLst) {
	for _, v := range cLst {
		if v.alias == s {
			c = v.structFieldLst
		}
	}
	return
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
	collectionNamesMap collectionLst
	aimer
	writter
	reader
	Err error
}

func newHandle() *Handle {
	d := &Handle{}
	return d
}
