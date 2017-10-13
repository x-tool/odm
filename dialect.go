package odm

import "reflect"

// database table
type Table struct {
	Name    string
	Type    reflect.Type
	colsMap map[string][]*Col
	Indexes map[string]*Index
	Created string
	Updated string
	Deleted string
	Version string
}

// config Struct
type ConnectionConfig struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	DatabaseName string
	Database     string
	TLs          bool
}

type Index string

type Dialect interface {
	Init(ConnectionConfig) Dialect
	// Conn() (Conn, error)
	GetColNames(db *Database) ([]string, error)
	SwitchType(string) string
	syncCol(*Col)
	// base handel
	Insert(doc *ODM) (interface{}, error)
	Update(doc *ODM) (interface{}, error)
	Delete(doc *ODM) (interface{}, error)
	Query(doc *ODM) (interface{}, error)
	Session() *Session
}

func initDialect(c ConnectionConfig) (d Dialect) {
	switch c.Database {
	case "postgresql":
		fallthrough
	default:
		_d := new(dialectpostgre)
		d = _d.Init(c)
		return d
	}
}

// type Conn interface {
// 	Open(sql string) error
// 	// Close()
// 	// Begin()
// }

type Session struct{}
type Exec interface{}
type Result interface{}
