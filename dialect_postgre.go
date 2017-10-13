package odm

import (
	"errors"
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
type postgreSimpleResult [][][]byte
type postgreConn struct {
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
func (d *dialectpostgre) GetColNames(db *Database) (ColNames []string, err error) {
	var results postgreSimpleResult
	// err = d.Open("SELECT tablename,tableowner FROM pg_tables WHERE schemaname='public'", results)
	err = d.Open("SELECT key,id FROM doc", results)
	log.Print(results)
	return ColNames, err
}

func (d *dialectpostgre) syncCol(col *Col) {
	var sql string
	var colFields string
	colName := col.name
	fieldLst := col.Doc.getRootDetails()
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
		if i == (fieldsNum - 1) {
			fieldPg = v.Name + " " + v.DBType
		} else {
			fieldPg = v.Name + " " + v.DBType + ",\n"
		}

		colFields = colFields + fieldPg
	}
	// make sql
	sql = fmt.Sprintf("CREATE TABLE %s ( %s ) ", colName, colFields)
	err := d.OpenWithODM(sql, nil)
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
	err = d.OpenWithODM(rawsql, result)
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

func (d *dialectpostgre) newConn() (*postgreConn, error) {
	conn, err := pgx.Connect(d.pgxConfig)
	var c postgreConn
	c.conn = conn
	return &c, err
}

func (d *dialectpostgre) Open(sql string, result interface{}) (results []interface{}, err error) {
	_conn, err := d.newConn()
	if err != nil {
		return nil, err
	}
	pgConn := _conn.conn
	log.Print(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()
	resultT := reflect.TypeOf(result)
	if resultT.Kind() != reflect.Struct {
		return nil, errors.New("ss")
	}
	resultSlice := reflect.MakeSlice(resultT, 1, 1)
	for rows.Next() {
		newResult := reflect.New(resultT)
		err := rows.Scan(newResult)
		log.Print(newResult)

		if err != nil {
			break
		}
	}
	return nil, err

}
func (d *dialectpostgre) OpenUseODM(sql string, result *ODM) (err error) {
	_conn, err := d.newConn()
	if err != nil {
		return err
	}
	pgConn := _conn.conn
	log.Print(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()

	for rows.Next() {
		_, err := rows.Values()
		// pg_ByteToValue(byteLst, result)
		if err != nil {
			break
		}
	}
	return err

}
