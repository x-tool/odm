package core

func (c *Col) GetRootDetails() (lst docFieldLst) {
	return c.doc.GetRootFields()
}
