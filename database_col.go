package xodm

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/zgr126/x-tool"
)

func (d *Database) newdatabaseCol(col ColInterface) error {
	p := reflect.ValueOf(col)
	v := p.Elem()
	t := v.Type()
	if v.Kind() == reflect.Struct {
		col := new(Col)
		col.Name = t.Name()
		colFieldNum := v.NumField()
		// make ColLst in a col
		for i := 0; i < colFieldNum; i++ {
			field := t.Field(i)
			FieldName := field.Name
			FieldTag := field.Tag.Get(tagName)
			FieldType = d.SwitchType(FieldTag)
			if FieldTag == "" {
				continue
			}
			col.detailLst = append(col.detailLst, &ColDetail{
				Name: FieldName,
				DetailType: 
			})
		}
		d.mergeCol(colName, colFieldLst)
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+v.Kind().String()+"!,Type should be Struct"))
	}
	return nil
}

func (d *Database) mergeCol(colName string, fieldLst []fieldStruct) {
	conn, err := d.Conn()
	if err != nil {
		tool.Panic("DB", err)
	}
	var sql string
	var colFields string
	fieldsNum := len(fieldLst)

	//output field name and typestr in colFields
	for i, v := range fieldLst {
		var fieldPg string
		// only one field abord ","
		if fieldsNum == 1 {
			fieldPg = v.name + " " + v.fieldTypeStr
			colFields = colFields + fieldPg
			break
		}
		// first and last rows abord tab and ","
		if i == 0 {
			fieldPg = v.name + " " + v.fieldTypeStr + ",\n"
		} else if i == (fieldsNum - 1) {
			fieldPg = "\t\t" + v.name + " " + v.fieldTypeStr
		} else {
			fieldPg = "\t\t" + v.name + " " + v.fieldTypeStr + ",\n"
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
