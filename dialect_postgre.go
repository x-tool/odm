package odm

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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
	config    ConnectionConfig
	pgxConfig pgx.ConnConfig
}
type postgredb struct {
	conn *pgx.Conn
}

func (d *dialectpostgre) Init(c ConnectionConfig) Dialect {
	d.config = c
	d.pgxConfig = pgx.ConnConfig{
		Host:     d.config.Host,
		Port:     uint16(d.config.Port),
		Database: d.config.DatabaseName,
		User:     d.config.User,
		Password: d.config.Passwd,
	}
	return d
}

func pg_valueToString(v reflect.Value) (r string) {
	_type := v.Kind().String()
	switch _type {
	case "string":
		r = "'" + v.String() + "'"
	case "int":
		r = strconv.FormatInt(v.Int(), 10)
	default:
		r = v.String()
	}
	return r
}

func (d *dialectpostgre) SwitchType(s string) string {
	return typeMap[s]
}
func (d *dialectpostgre) GetTables(db *Database) ([]string, error) {
	var tablesName []string
	conn, _ := d.Conn()
	type r struct {
		tablesName string
	}
	var rLst []r
	err := conn.Open("SELECT tablename FROM pg_tables WHERE schemaname='public'", rLst)
	log.Print(rLst)
	return tablesName, err
}

func (d *dialectpostgre) syncCol(col *Col) {
	conn, err := d.Conn()
	if err != nil {
		tool.Panic("DB", err)
	}
	var sql string
	var colFields string
	colName := col.name
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
	err = conn.Open(sql, nil)
	if err != nil {
		tool.Panic("DB", err)
	}
}

func (d *dialectpostgre) Session() *Session {
	return new(Session)
}
func (d *dialectpostgre) Insert(doc *ODM) (result interface{}, err error) {
	var nameLst, valueLst []string
	rootFields := doc.Result.getRootFields()
	for _, v := range rootFields {
		nameLst = append(nameLst, v.name)
		valueLst = append(valueLst, pg_valueToString(v.value))
	}
	nameLstStr := strings.Join(nameLst, ",")
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName ($typeLst) VALUES ($valueLst)"
	rawsql := tool.ReplaceStrings(sql, []string{
		"$colName", doc.Col.name,
		"$typeLst", nameLstStr,
		"$valueLst", valueLstStr,
	})
	conn, _ := d.Conn()
	err = conn.Open(rawsql, result)
	log.Println(result)
	return
}
func (d *dialectpostgre) Update(doc *ODM) (r interface{}, err error) {
	return
}
func (d *dialectpostgre) Delete(doc *ODM) (r interface{}, err error) {
	return
}
func (d *dialectpostgre) Query(doc *ODM) (r interface{}, err error) {
	return
}

func (d *dialectpostgre) Conn() (Conn, error) {
	conn, err := pgx.Connect(d.pgxConfig)
	var c postgreConn
	c.conn = conn
	return &c, err
}

type postgreConn struct {
	conn *pgx.Conn
}

func (p *postgreConn) Open(s string, result interface{}) (err error) {
	sql := s
	log.Print(sql)
	rows, err := p.conn.Query(sql)
	defer p.conn.Close()
	var rLst [][]interface{}
	for rows.Next() {
		// var tableName string
		byteLst, err := rows.Values()
		rLst = append(rLst, byteLst)
		// log.Print(s)
		if err != nil {
			break
		}
	}
	log.Println(rLst)
	return err
}

// func (p *postgreConn)Q(sql string)(r interface{},err error){

// }
// func (p *postgredb) Close(ql string) (result []byte, err error) {
// 	result = []byte{}
// 	return result, err
// }
