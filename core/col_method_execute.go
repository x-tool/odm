package core

func (c *Col) Insert(i interface{}) (err error) {
	return c.db.Insert(i)
}
