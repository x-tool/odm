package dialect

import (
	"github.com/x-tool/odm/module"
	"github.com/x-tool/odm/module/dialect/model"
)

type Dialect = model.Dialect

func NewDialect(c module.ConnectConfig) (d Dialect) {
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
