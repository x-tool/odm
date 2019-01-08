package odm

import (
	"testing"

	"github.com/x-tool/odm/core"
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

func Test_Insert(t *testing.T) {
	db := connection()
	testdata := new(myDoc)
	testdata.Id = 1
	errInsert := db.Insert(testdata)
	t.Log(errInsert)
}

func Test_Query(t *testing.T) {
	db := connection()
	testdata := new(myDoc)
	errInsert := db.Query(testdata).Where("(name = ?) and createTime = ", "LiLei").Get()
	_ = db.Query(nil)
	// col := db.GetCol(new(myDoc))
	// _, err := col.data(testdata)
	// log.Print(errInsert)
	// col.Key(testdata.Key).Delete()
}

func connection() (db *core.Database) {
	connectionConf := Connect{
		Host:         "127.0.0.1",
		Port:         5432,
		User:         "postgres",
		Passwd:       "zxczxczxc",
		Database:     "postgresql",
		DatabaseName: "x",
	}
	odm := New(connectionConf)
	db = odm.Database()
	db.RegisterCols(new(myDoc))
	db.SyncCols()
	return
}

// func Test_formatType(t *testing.T) {
// 	var a string
// 	s := reflect.TypeOf(a)
// 	str := formatTypeToString(&s)
// 	t.Log(str)
// }
