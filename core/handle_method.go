package core

import (
	"errors"
	"fmt"
	"reflect"
)

func (h *Handle) GetDBName() (s string, e error) {
	if len(h.handleCols) == 0 {
		return "", errors.New("no col in handle")
	} else {
		return h.handleCols[0].col.db.name, nil
	}
}

func (h *Handle) getStructByStr(s string) (o *odmStruct, err error) {
	for k, v := range h.alias {
		if s == k {
			return v, nil
		}
	}
	for k, v := range h.db.mapStructs {
		if s == k {
			return v, nil
		}
	}
	return nil, fmt.Errorf("can't get struct use string: %v", s)
}

func (h *Handle) getColByStr(s string) (o *Col, err error) {
	name := s
	for k, v := range h.alias {
		if s == k {
			name = v.name
		}
	}
	for k, v := range h.db.mapCols {
		if name == k {
			return v, nil
		}
	}
	return nil, fmt.Errorf("can't get Col use string: %v", s)
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
