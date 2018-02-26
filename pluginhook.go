package odm

// config Struct
type ConnectConfig struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	DatabaseName string
	Database     string
	TLs          bool
}

type Dialect interface {
	Init(ConnectConfig) Dialect
	// Conn() (Conn, error)
	GetColNames() ([]string, error)
	SwitchType(string) string
	syncCol(*odm.Col)
	// base handel
	Insert(*odm.Handle) error
	Update(*odm.Handle) error
	Delete(*odm.Handle) error
	Query(*odm.Handle) (interface{}, error)
	LogSql(string)
	Session() *Session
}

func NewDialect(c ConnectConfig) (d Dialect) {
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

type DocModer interface {
	config()
}
