package odm

import (
	"errors"
	"fmt"
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
	// case "time":
	// 	if v
	default:
		r = v.String()
	}
	return r
}

func (d *dialectpostgre) SwitchType(s string) string {
	return typeMap[s]
}
func (d *dialectpostgre) GetColNames(db *Database) (ColNames []string, err error) {
	type _result struct {
		Tablename  string
		Tableowner string
	}
	var results []_result
	err = d.Open("SELECT tablename,tableowner FROM pg_tables WHERE schemaname='public'", &results)
	for _, v := range results {
		ColNames = append(ColNames, v.Tablename)
	}
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
func (d *dialectpostgre) Insert(doc *ODM) (err error) {
	var nameLst, valueLst []string
	rootFields := doc.Result.getRootFields()
	rootFields = doc.selectValidFields(rootFields)
	for _, v := range rootFields {
		nameLst = append(nameLst, v.DocField.Name)
		valueLst = append(valueLst, pg_valueToString(v.value))
	}
	nameLstStr := strings.Join(nameLst, ",")
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName ($typeLst) VALUES ($valueLst) RETURNING *"
	rawsql := tool.ReplaceStrings(sql, []string{
		"$colName", doc.Col.name,
		"$typeLst", nameLstStr,
		"$valueLst", valueLstStr,
	})
	err = d.OpenWithODM(rawsql, nil)
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
func (d *dialectpostgre) LogSql(sql string) {
	tool.Console.Log("XODM", sql)
}
func (d *dialectpostgre) Open(sql string, results interface{}) (err error) {
	_conn, err := d.newConn()
	if err != nil {
		return err
	}
	pgConn := _conn.conn
	d.LogSql(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()

	resultV := reflect.ValueOf(results)
	resultVElem := reflect.Indirect(resultV)
	resultT := resultVElem.Type()
	if resultT.Kind() != reflect.Slice {
		return errors.New("result type should be slice, Not Be " + resultT.Kind().String())
	} else {
		resultItemT := resultT.Elem()
		_tempResultItemLst := reflect.New(reflect.SliceOf(resultItemT))
		tempResultItemLst := reflect.Indirect(_tempResultItemLst)
		for rows.Next() {
			newResult := reflect.Indirect(reflect.New(resultItemT))
			var resultSlicePtr []interface{}
			for i := 0; i < newResult.NumField(); i++ {
				newResultField := newResult.Field(i).Addr().Interface()
				resultSlicePtr = append(resultSlicePtr, newResultField)
			}
			err := rows.Scan(resultSlicePtr...)
			tempResultItemLst.Set(reflect.Append(tempResultItemLst, newResult))
			if err != nil {
				break
			}
		}
		resultVElem.Set(tempResultItemLst)
		return err
	}
}
func (d *dialectpostgre) OpenWithODM(sql string, result *ODM) (err error) {
	_conn, err := d.newConn()
	if err != nil {
		return err
	}
	pgConn := _conn.conn
	d.LogSql(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()

	resultV := *(result.R)
	resultT := resultV.Type()
	if resultT.Kind() != reflect.Slice {
		return errors.New("result type should be slice, Not Be " + resultT.Kind().String())
	} else {
		resultItemT := resultT.Elem()
		_tempResultItemLst := reflect.New(reflect.SliceOf(resultItemT))
		tempResultItemLst := reflect.Indirect(_tempResultItemLst)
		for rows.Next() {
			newResult := reflect.Indirect(reflect.New(resultItemT))
			var resultSlicePtr []interface{}
			for i := 0; i < newResult.NumField(); i++ {
				newResultField := newResult.Field(i).Addr().Interface()
				resultSlicePtr = append(resultSlicePtr, newResultField)
			}
			err := rows.Scan(resultSlicePtr...)
			tempResultItemLst.Set(reflect.Append(tempResultItemLst, newResult))
			if err != nil {
				break
			}
		}
		resultV.Set(tempResultItemLst)
		return err
	}

}
