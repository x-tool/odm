package xodm

import (
	"fmt"
	"log"

	"github.com/x-tool/tool"
)

func (d *Database) syncCol(colI ColInterface) {
	col := NewCol(colI)
	syncCol(col)
}

func (d *Database) newCol(colName string, fieldLst []fieldStruct) {
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
