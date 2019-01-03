package odm

import (
	"testing"
)

type myDocLabel struct {
	// Name       string
	Label      string `xodm:"notnull default:你好"`
	LabelName  string `xodm:"default:"`
	Labelfeild string
}
type myDoc struct {
	NormalCol
	Name string `xodm:"notnull"`
	Id   int
	myDocLabel
}

func (m *myDoc) ColName() string {
	return "doc"
}
func Test_connection(t *testing.T) {
	connectionConf := Connect{
		Host:         "127.0.0.1",
		Port:         5432,
		User:         "postgres",
		Passwd:       "zxczxczxc",
		Database:     "postgresql",
		DatabaseName: "x",
	}
	odm := New(connectionConf)
	db := odm.Database()
	db.RegisterCols(new(myDoc))
	db.SyncCols()

	testdata := new(myDoc)
	// testdata.Name = "LiLei"
	testdata.Id = 1
	// errInsert := db.Insert(testdata)
	// errInsert := db.Delete(testdata).Where("name = LiLei")
	_ = db.Query(nil)
	// col := db.GetCol(new(myDoc))
	// _, err := col.data(testdata)
	// log.Print(errInsert)
	// col.Key(testdata.Key).Delete()
	// t.Log(errInsert)
}

// func Test_formatType(t *testing.T) {
// 	var a string
// 	s := reflect.TypeOf(a)
// 	str := formatTypeToString(&s)
// 	t.Log(str)
// }
