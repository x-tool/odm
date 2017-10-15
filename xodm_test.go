package odm

import (
	"reflect"
	"testing"
)

type myDocLabel struct {
	// Name       string
	Label      string
	labelName  string
	labelfeild string
}
type myDoc struct {
	NormalCol
	Name   string
	Id     int
	Detail *myDocLabel `odm:"extend"`
}

func (m *myDoc) ColName() string {
	return "doc"
}
func Test_connection(t *testing.T) {
	connectionConf := ConnectionConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Passwd:   "zxczxczxc",
		Database: "postgresql",
	}
	client := NewClient("postgresql", connectionConf)
	db := client.Database("x")

	db.SyncCols(new(myDoc))
	col := db.GetCol(new(myDoc))
	testInsert := new(myDoc)
	testInsert.Name = "haha,I get"
	_, err := col.Insert(testInsert)
	t.Log(testInsert)
	t.Log(err)
}

func Test_formatType(t *testing.T) {
	var a string
	s := reflect.TypeOf(a)
	str := formatTypeToString(&s)
	t.Log(str)
}
