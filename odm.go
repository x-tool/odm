package odm

func NewClient( conf ConnectionConfig) *client {
	var c client
	c.dbName = dbName
	c.config = conf
	return &c
}