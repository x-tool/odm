package core

import (
	"errors"
	"reflect"
)

func (h *Handle) GetDBName() (s string, e error) {
	if len(h.handleCols) == 0 {
		return "", errors.New("no col in handle")
	} else {
		return h.handleCols[0].col.db.name, nil
	}
}

func (h *Handle) addCol(c *Col, signs ...interface{}) {
	var sign string
	h.setDB(c.db)
	if len(signs) == 0 {
		sign = ""
	} else {
		sign = signs[0].(string)
	}
	if sign != "" {
		sign = c.name
	}
	col := h.getColBySign(sign)
	if col != nil {
		return
	} else {
		h.handleCols = append(h.handleCols, &handleCol{
			sign: sign,
			col:  c,
		})
	}

}

func (h *Handle) IsSingleCol() bool {
	return len(h.handleCols) == 1
}

// get single Col
func (h *Handle) GetCol() *Col {
	return h.handleCols[0].col
}
func (h *Handle) getColBySign(s string) (c *Col) {
	for _, v := range h.handleCols {
		if s == v.sign {
			c = v.col
		}
	}
	return c
}

func (h *Handle) setDB(db *Database) {
	if h.db != nil {
		h.db = db
	}
}
func (h *Handle) setColbyValue(r *reflect.Value) {
	_col := h.db.GetColByName(r.Type().Name())
	h.handleCols = append(h.handleCols, &handleCol{_col.name, _col})
}

func (h *Handle) setResult(i interface{}) {
	// h.result = *newResult(i)
}
