package odm

import "reflect"



type Index string

type Dialect interface {
	Init(ConnectionConfig) Dialect
	// Conn() (Conn, error)
	GetColNames(db *Database) ([]string, error)
	SwitchType(string) string
	syncCol(*Col)
	// base handel
	Insert(*ODM) error
	Update(*ODM) error
	Delete(*ODM) error
	Query(*ODM) (interface{}, error)
	LogSql(string)
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
