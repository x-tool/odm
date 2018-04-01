package core

type doc struct {
	col     *Col
	docMode structField
	odmStruct
}

type docLst []*doc

func newDoc(c *Col, i interface{}) *doc {
	_doc := new(doc)
	_doc.col = c
	_doc.odmStruct = *newOdmStruct(i)
	return _doc
}
