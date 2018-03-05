package odm

import (
	"testing"
)

type myDocLabel struct {
	// Name       string
	Label      string
	LabelName  string
	Labelfeild string
}
type myDoc struct {
	NormalCol
	Name       string
	Id         int
	myDocLabel `Handle:"extend"`
}

func (m *myDoc) ColName() string {
	return "doc"
}
func Test_connection(t *testing.T) {
	connectionConf := Connect{
		Host:         "127.0.0.1",
		Port:         5432,
		User:         "postgre",
		Passwd:       "zxczxczxc",
		Database:     "postgresql",
		DatabaseName: "x",
	}
	odm := New(connectionConf)
	db := odm.Database(nil)
	db.RegisterCols(new(myDoc))
	db.SyncCols()
	// col := db.GetCol(new(myDoc))
	// testInsert := new(myDoc)
	// testInsert.Name = "haha,I get"
	// testInsert.Id = 1
	// _, err := col.Insert(testInsert)
	// log.Print(testInsert)
	// col.Key(testInsert.Key).Delete()
	// t.Log(err)
}

// func Test_formatType(t *testing.T) {
// 	var a string
// 	s := reflect.TypeOf(a)
// 	str := formatTypeToString(&s)
// 	t.Log(str)
// }
