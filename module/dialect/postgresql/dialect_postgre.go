package Handle

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/x-tool/odm/exporti"
	"github.com/x-tool/tool"
)

const (
	pg_timeFormat = "2006-01-02 03:04:05"
)

var typeMap = map[string]string{
	"int":     "int",
	"float64": "float8",
	// "text":   "text",
	"[]byte": "bytea",
	"time":   "timestamp",
	"array":  "json",
	"bool":   "json",
	"string": "text",
	"struct": "json",
	"Others": "text",
}

type dialectpostgre struct {
	config    exporti.ConnectionConfig
	pgxConfig pgx.ConnConfig
}

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

func pg_valueToString(v *queryRootField) (r string) {
	value := v.value
	switch v.DocField.Type {
	case "string":
		r = "'" + value.String() + "'"
	case "int":
		r = strconv.FormatInt(value.Int(), 10)
	case "time":
		if _v, ok := value.Interface().(time.Time); ok {
			r = "'" + _v.Format(pg_timeFormat) + "'"
		}
	default:
		r = value.String()
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
	err := d.Open(sql, nil)
	if err != nil {
		tool.Panic("DB", err)
	}
}

func (d *dialectpostgre) Session() *Session {
	return new(Session)
}
func (d *dialectpostgre) Insert(doc *Handle) (err error) {
	var nameLst, valueLst []string
	rootFields := doc.Query.getRootFields()
	rootFields = doc.selectValidFields(rootFields)
	for _, v := range rootFields {
		nameLst = append(nameLst, v.DocField.Name)
		valueLst = append(valueLst, pg_valueToString(v))
	}
	nameLstStr := strings.Join(nameLst, ",")
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName ($typeLst) VALUES ($valueLst) RETURNING *"
	rawsql := tool.ReplaceStrings(sql, []string{
		"$colName", doc.Col.name,
		"$typeLst", nameLstStr,
		"$valueLst", valueLstStr,
	})
	err = d.OpenWithHandle(rawsql, doc)
	return
}
func (d *dialectpostgre) Update(doc *Handle) (err error) {
	return
}
func (d *dialectpostgre) Delete(doc *Handle) (err error) {
	return
}
func (d *dialectpostgre) Query(doc *Handle) (r interface{}, err error) {
	return
}

func (d *dialectpostgre) newConn() (*postgreConn, error) {
	conn, err := pgx.Connect(d.pgxConfig)
	var c postgreConn
	c.conn = conn
	return &c, err
}
func (d *dialectpostgre) LogSql(sql string) {
	tool.Console.LogWithLabel("XHandle", sql)
}

func pg_formatQL(o *Handle) (s string) {
	var queryStr string
	var resultStr string
	for i, v := range o.Result.resultFieldLst {
		var _resultStr string
		vRootField := v.getRootFieldDB()
		if i == 0 && vRootField.Type == "struct" {
			jsonStr := vRootField.Name
			for i, _v := range v.getDependLstDB() {
				if i == 0 {
					continue
				} else {
					jsonStr = jsonStr + "->'" + _v.Name + "'"
				}

			}
			_resultStr = jsonStr
		} else {
			_resultStr = vRootField.Name
		}

		if i == 0 {
			resultStr = _resultStr
		} else {
			resultStr = resultStr + "," + _resultStr
		}
	}

	for _, v := range o.Query.queryLst {
		var _resultStr string
		vRootField := v.getRootFieldDB()
		if i == 0 && vRootField.Type == "struct" {
			jsonStr := vRootField.Name
			for i, _v := range v.getDependLstDB() {
				if i == 0 {
					continue
				} else {
					jsonStr = jsonStr + "->'" + _v.Name + "'"
				}

			}
			_resultStr = jsonStr
		} else {
			_resultStr = vRootField.Name
		}

		if i == 0 {
			resultStr = _resultStr
		} else {
			resultStr = resultStr + "," + _resultStr
		}
	}
	s = "SELECT " + queryStr + " FROM " + o.colName() + "WHERE " + resultStr
	return
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
	if results == nil {
		return err
	}
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
func (d *dialectpostgre) OpenWithHandle(sql string, Handle *Handle) (err error) {
	_conn, err := d.newConn()
	if err != nil {
		return err
	}
	pgConn := _conn.conn
	d.LogSql(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()

	resultV := *Handle.R
	resultT := resultV.Type()
	if resultT.Kind() == reflect.Slice {
		for rows.Next() {
			resultItemV := Handle.Result.newResultItem()
			var resultSlicePtr []interface{}
			for _, v := range Handle.Result.getResultRootItemFieldAddr(resultItemV) {
				resultSlicePtr = append(resultSlicePtr, (v).Interface())
			}
			err = rows.Scan(resultSlicePtr...)
			resultV.Set(reflect.Append(resultV, *resultItemV))
			if err != nil {
				break
			}
		}
		return err
	} else {
		return nil
	}

}
