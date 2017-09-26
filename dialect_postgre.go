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
func (d *dialectpostgre) Insert(doc *Doc) (r interface{}, err error) {
	var typeLst, valueLst []string
	rootFields := doc.getRootfields()
	for _, v := range rootFields {
		typeLst = append(typeLst, v.DBtypeName)
		valueLst = append(valueLst, pg_valueToString(v.value))
	}
	typeLstStr := strings.Join(typeLst, ",")
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName ($typeLst) VALUES ($valueLst)"
	rawsql := tool.ReplaceStrings(sql, map[string]string{
		"$colName":  doc.Col.Name,
		"$typeLst":  typeLstStr,
		"$valueLst": valueLstStr,
	})
	conn, _ := d.Conn()
	result, err := conn.Open(rawsql)
	log.Println(rawsql)
	log.Println(result)
	log.Println(err)
	return
}
func (d *dialectpostgre) Update(doc *Doc) (r interface{}, err error) {
	return
}
func (d *dialectpostgre) Delete(doc *Doc) (r interface{}, err error) {
	return
}
func (d *dialectpostgre) Query(doc *Doc) (r interface{}, err error) {
	return
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
	rows, err := p.conn.Query(sql)
	defer p.conn.Close()
	log.Print(rows)
	for rows.Next() {
		// var tableName string
		s, err := rows.Values()
		if err != nil {
			break
		}
		log.Println(s)
	}
	return result, err
}

// func (p *postgreConn)Q(sql string)(r interface{},err error){

// }
// func (p *postgredb) Close(ql string) (result []byte, err error) {
// 	result = []byte{}
// 	return result, err
// }
