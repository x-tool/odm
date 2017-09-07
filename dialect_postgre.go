package xodm

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/x-tool/tool"
)

var typeMap = map[string]string{
	"int":    "int",
	"float":  "float8",
	"text":   "text",
	"[]byte": "bytea",
	"time":   "date",
	"array":  "json",
	"bool":   "json",
	"string": "text",
	"struct": "json",
	"Others": "text",
}

type dialectpostgre struct {
	config ConnectionConfig
}
type postgredb struct {
	conn *pgx.Conn
}

func (d *dialectpostgre) Init(c ConnectionConfig) Dialect {
	d.config = c
	return d
}

func (d *dialectpostgre) SwitchType(s string) string {
	return typeMap[s]
}
func (d *dialectpostgre) GetTables() ([]string, error) {
	var tablesName []string
	conn, _ := d.Conn()
	_, err := conn.Open("SELECT tablename FROM pg_tables WHERE schemaname='public'")
	return tablesName, err
}

func (d *dialectpostgre) syncCol(col *Col) {
	conn, err := d.Conn()
	if err != nil {
		tool.Panic("DB", err)
	}
	var sql string
	var colFields string
	colName := col.Name
	fieldLst := col.getRootDetails()
	fieldsNum := len(fieldLst)

	//output field name and typestr in colFields
	for i, v := range fieldLst {
		var fieldPg string
		// only one field abord ","
		if fieldsNum == 1 {
			fieldPg = v.Name + " " + v.DBType
			colFields = colFields + fieldPg
			break
		}
		// first and last rows abord tab and ","
		if i == 0 {
			fieldPg = v.Name + " " + v.DBType + ",\n"
		} else if i == (fieldsNum - 1) {
			fieldPg = "\t\t" + v.Name + " " + v.DBType
		} else {
			fieldPg = "\t\t" + v.Name + " " + v.DBType + ",\n"
		}

		colFields = colFields + fieldPg
	}
	// make sql
	sql = fmt.Sprintf(`
		CREATE TABLE %s
		(
			%s
		)
	`, colName, colFields)
	log.Print(sql)
	_, err = conn.Open(sql)
	if err != nil {
		tool.Panic("DB", err)
	}
}

func (d *dialectpostgre) Session() *Session {
	return new(Session)
}
func (d *dialectpostgre) Insert(result interface{}, o *docStruct) {

}
func (d *dialectpostgre) Update(result interface{}, o *docStruct) {

}
func (d *dialectpostgre) Delete(result interface{}, o *docStruct) {

}
func (d *dialectpostgre) Query(result interface{}, o *docStruct) {

}

type postgreConn struct {
	conn *pgx.Conn
}

func (d *dialectpostgre) Conn() (Conn, error) {
	log.Print(d.config)
	var pgxConf pgx.ConnConfig
	pgxConf = pgx.ConnConfig{
		Host:     d.config.Host,
		Port:     uint16(d.config.Port),
		Database: d.config.DatabaseName,
		User:     d.config.User,
		Password: d.config.Passwd,
	}
	conn, err := pgx.Connect(pgxConf)
	var c postgreConn
	c.conn = conn
	return &c, err
}

func (p *postgreConn) Open(s ...Exec) (result Result, err error) {
	var sql string
	if len(s) == 1 {
		s, ok := s[0].(string)
		if ok {
			sql = s
		}
	}
	re, err := p.conn.Exec(sql)
	defer p.conn.Close()
	log.Print(re)
	return result, err
}

// func (p *postgredb) Close(ql string) (result []byte, err error) {
// 	result = []byte{}
// 	return result, err
// }
