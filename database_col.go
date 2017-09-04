package xodm

import (
	"fmt"
	"log"

	"github.com/x-tool/tool"
)

func (d *Database) syncCol(colI ColInterface) {
	col := NewCol(d, colI)
	syncColToDB(d, col)
}

func syncColToDB(d *Database, col *Col) {
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
