package xodm

import "testing"

type myDocLabel struct {
	// Name       string
	Label      string
	labelName  string
	labelfeild string
}
type myDoc struct {
	Normal
	Name   string
	Id     int
	Detail *myDocLabel `xodm:"extend"`
}

func (m *myDoc) ColName() string {
	return "doc"
}
func Test_connection(t *testing.T) {
	connectionConf := ConnectionConfig{
		Host:   "127.0.0.1",
		Port:   5432,
		User:   "postgres",
		Passwd: "zxczxc",
	}
	client := NewClient("postgresql", connectionConf)
	db := client.Database("x")

	db.SyncCols(new(myDoc))
	var r interface{}
	db.Insert(r, new(myDoc))
	t.Log(db)
}
