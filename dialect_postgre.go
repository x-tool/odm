package xodm

import (
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx"
	"github.com/x-tool/tool"
)

var typeMap = map[string]string{
	"int":     "int",
	"float64": "float8",
	// "text":   "text",
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

func pg_valueToString(i interface{}) (r string) {
	switch i.(type) {
	case string:
		r = i.(string)
	case int:
		r = i.(string)
	case float64:
		r = i.(string)
		// case byte:
		// 	r = string()
	}
	return
}

func (d *dialectpostgre) SwitchType(s string) string {
	return typeMap[s]
}
func (d *dialectpostgre) GetTables() ([]string, error) {
	var tablesName []string
	conn, _ := d.Conn()
	r, err := conn.Open("SELECT tablename FROM pg_tables WHERE schemaname='public'")
	log.Println(r, err)
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
	fieldLst := col.OriginDocs.getRootDetails()
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
func (d *dialectpostgre) Insert(c *Col, result interface{}, i interface{}) {
	var typeLst, valueLst []string
	rootFields := c.getRootfields(i)
	for _, v := range rootFields {
		typeLst = append(typeLst, v.DBtypeName)
		valueLst = append(valueLst, pg_valueToString(v.value))
	}
	typeLstStr := strings.Join(typeLst, ",")
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName ($typeLst) VALUES ($valueLst)"
	rawsql := tool.ReplaceStrings(sql, map[string]string{
		"$colName":  c.Name,
		"$typeLst":  typeLstStr,
		"$valueLst": valueLstStr,
	})
	conn, _ := d.Conn()
	r, err := conn.Open(rawsql)
	log.Println(r)
	log.Println(err)
}
func (d *dialectpostgre) Update(c *Col, result interface{}, i interface{}) {

}
func (d *dialectpostgre) Delete(c *Col, result interface{}, i interface{}) {

}
func (d *dialectpostgre) Query(c *Col, result interface{}, i interface{}) {

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
	re, err := p.conn.Query(sql)
	defer p.conn.Close()
	log.Print(re)
	b, _ := re.Values()
	log.Print(b)
	return result, err
}

// func (p *postgredb) Close(ql string) (result []byte, err error) {
// 	result = []byte{}
// 	return result, err
// }
