package postgresql

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx"
	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
	"github.com/x-tool/tool"
)

const (
	pg_timeFormat = "2006-01-02 03:04:05"
)

var typeMap = map[core.Kind]string{
	core.Int:   "int",
	core.Float: "float8",
	// "text":   "text",
	core.Byte:   "bytea",
	core.Time:   "timestamp",
	core.Array:  "json",
	core.Bool:   "json",
	core.String: "text",
	core.Struct: "json",
}

func kindToString(k core.Kind) (s string) {
	return typeMap[k]
}

func valueToString(field *core.StructField, value *reflect.Value) (str string) {
	var isZero = !value.IsValid()
	// if value is zero value
	if isZero {
		if !field.NotNull() {
			return "null"
		} else {
			// has, value := field.Default()
			// if has {
			// 	value = reflect.ValueOf
			// }
		}
	}
	var kind = field.Kind()
	// postgre type handle
	switch kind {
	case core.Time:
		var t time.Time
		if isZero {
			_t := new(time.Time)
			t = *_t
		} else {
			t = value.Interface().(time.Time)
		}
		str = t.Format("2006-01-02 15:04:05")

		// default use core default process mode
		// default:
		// 	str = core.ValueToString(value)
	}
	if kind == core.Int {
		if str == "" {
			str = "0"
		}
		str = str
	} else {
		str = "'" + str + "'"
	}
	return
}

type postgreConn struct {
	conn *pgx.Conn
}

func New() core.Dialect {
	return new(dialectpostgre)
}

type dialectpostgre struct {
	config          *client.ConnectConfig
	pgxConfig       pgx.ConnConfig
	historyColNames []string
}

func (d *dialectpostgre) Init(c *client.Client) {
	d.config = c.GetConnectConfig()
	d.pgxConfig = pgx.ConnConfig{
		Host:     d.config.Host,
		Port:     uint16(d.config.Port),
		Database: d.config.DatabaseName,
		User:     d.config.User,
		Password: d.config.Passwd,
	}
	// return d
}

// func pg_valueToString(v *queryRootField) (r string) {
// 	value := v.value
// 	switch v.DocField.Type {
// 	case "string":
// 		r = "'" + value.String() + "'"
// 	case "int":
// 		r = strconv.FormatInt(value.Int(), 10)
// 	case "time":
// 		if _v, ok := value.Interface().(time.Time); ok {
// 			r = "'" + _v.Format(pg_timeFormat) + "'"
// 		}
// 	default:
// 		r = value.String()
// 	}
// 	return r
// }

// func (d *dialectpostgre) SwitchType(s string) string {
// 	return typeMap[s]
// }

func (d *dialectpostgre) GetColNames() (ColNames []string, err error) {
	type _result struct {
		Tablename  string
		Tableowner string
	}
	var results []_result
	err = d.Open("SELECT tablename,tableowner FROM pg_tables WHERE schemaname='public'", &results)
	for _, v := range results {
		d.historyColNames = append(d.historyColNames, v.Tablename)
	}
	return d.historyColNames, err
}

func (d *dialectpostgre) SyncCols(colLst core.ColLst) {
	var syncLock sync.WaitGroup
	for _, col := range colLst {
		syncLock.Add(1)
		go func(*core.Col) {
			defer syncLock.Done()
			var sql string
			var colFields string
			colName := col.Name()
			fieldLst := col.GetRootFields()
			var fieldStringLst []string
			//output field name and typestr in colFields
			for _, field := range fieldLst {
				fieldKind := kindToString(field.Kind())
				name := field.Name()
				// add field's constraints
				var constraints string
				if field.NotNull() {
					constraints = "NOT NULL"
				}
				// ***** if use 'if else' at here. field map become map[],can't get value if default has value,why???
				has, defaultValue := field.DefaultValue()
				if has {
					reflectV := reflect.ValueOf(defaultValue)
					constraints = fmt.Sprintf("%v Default %v", constraints, valueToString(field, &reflectV))
				}

				fieldStringLst = append(fieldStringLst, fmt.Sprintf("%v %v %v", name, fieldKind, constraints))
			}

			colFields = tool.JoinStringWithComma(fieldStringLst)
			// make sql
			sql = fmt.Sprintf("CREATE TABLE %s ( %s ) ", colName, colFields)
			err := d.Open(sql, nil)
			if err != nil {
				tool.Panic("DB", err)
			}
		}(col)
	}
	syncLock.Wait()
}

func (d *dialectpostgre) Session() *core.Session {
	return new(core.Session)
}
func (d *dialectpostgre) Insert(h *core.Handle) (err error) {
	var valueLst []string
	_col := h.GetCol()
	fields := _col.GetRootFields()
	_valueLst := _col.GetRootValues(h.GetWritterValue())
	if err != nil {
		return
	}
	for i := 0; i < len(fields); i++ {
		valueLst = append(valueLst, valueToString(fields[i], &_valueLst[i]))
	}
	valueLstStr := strings.Join(valueLst, ",")
	sql := "INSERT INTO $colName VALUES ($valueLst) RETURNING *"
	rawsql := tool.ReplaceStrings(sql, []string{
		"$colName", _col.Name(),
		"$valueLst", valueLstStr,
	})
	err = d.Open(rawsql, nil)
	// err = d.OpenWithHandle(rawsql, h)
	return
}
func (d *dialectpostgre) Update(doc *core.Handle) (err error) {
	return
}
func (d *dialectpostgre) Delete(doc *core.Handle) (err error) {
	return
}
func (d *dialectpostgre) Query(doc *core.Handle) (err error) {
	err = d.Open("select name, label from mydoc", []*[]byte{})
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

// func pg_formatQL(o *core.Handle) (s string) {
// 	var queryStr string
// 	var resultStr string
// 	for i, v := range o.Result.resultFieldLst {
// 		var _resultStr string
// 		vRootField := v.getRootFieldDB()
// 		if i == 0 && vRootField.Type == "struct" {
// 			jsonStr := vRootField.Name
// 			for i, _v := range v.getDependLstDB() {
// 				if i == 0 {
// 					continue
// 				} else {
// 					jsonStr = jsonStr + "->'" + _v.Name + "'"
// 				}

// 			}
// 			_resultStr = jsonStr
// 		} else {
// 			_resultStr = vRootField.Name
// 		}

// 		if i == 0 {
// 			resultStr = _resultStr
// 		} else {
// 			resultStr = resultStr + "," + _resultStr
// 		}
// 	}

// 	for _, v := range o.Query.queryLst {
// 		var _resultStr string
// 		vRootField := v.getRootFieldDB()
// 		if i == 0 && vRootField.Type == "struct" {
// 			jsonStr := vRootField.Name
// 			for i, _v := range v.getDependLstDB() {
// 				if i == 0 {
// 					continue
// 				} else {
// 					jsonStr = jsonStr + "->'" + _v.Name + "'"
// 				}

// 			}
// 			_resultStr = jsonStr
// 		} else {
// 			_resultStr = vRootField.Name
// 		}

// 		if i == 0 {
// 			resultStr = _resultStr
// 		} else {
// 			resultStr = resultStr + "," + _resultStr
// 		}
// 	}
// 	s = "SELECT " + queryStr + " FROM " + o.colName() + "WHERE " + resultStr
// 	return
// }

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
		for {
			if !rows.Next() {
				break
			}
			// newResult := reflect.Indirect(reflect.New(resultItemT))
			// var resultSlicePtr [][]byte
			// for i := 0; i < newResult.NumField(); i++ {
			// 	// newResultField := newResult.Field(i).Addr().Interface()
			// 	resultSlicePtr = append(resultSlicePtr, []byte{})
			// }
			var a []byte
			var b []byte
			err := rows.Scan(&a, &b)
			log.Println("asdf:", string(a), string(b))
			// tempResultItemLst.Set(reflect.Append(tempResultItemLst, newResult))
			if err != nil {
				break
			}
		}
		resultVElem.Set(tempResultItemLst)
		return err
	}
}

func (d *dialectpostgre) OpenWithHandle(sql string, h *core.Handle) (err error) {
	_conn, err := d.newConn()
	if err != nil {
		return err
	}
	pgConn := _conn.conn
	d.LogSql(sql)
	rows, err := pgConn.Query(sql)
	defer pgConn.Close()

	if h.Reader.IsComplex() {
		for rows.Next() {
			err = rows.Scan(h.Reader.Row().Addrs()...)
			if err != nil {
				return err
			}
		}
	} else {
		for rows.Next() {
			err = rows.Scan(h.Reader.Row().Addrs()...)
			if err != nil {
				return err
			}
		}
	}
	return err
}
