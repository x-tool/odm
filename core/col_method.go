package core

func (c *Col) GetRootFields() (lst docFieldLst) {
	return c.doc.GetRootFields()
}
