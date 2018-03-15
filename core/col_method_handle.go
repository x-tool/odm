package core

func (c *Col) Insert(i interface{}) *Handle {
	return newHandleByCol(c, nil)
}
