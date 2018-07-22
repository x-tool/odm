package core

import (
	"errors"
	"reflect"
)

func (h *Handle) GetDBName() (s string, e error) {
	if len(h.ColLst) == 0 {
		return "", errors.New("no col in handle")
	} else {
		return h.ColLst[0].db.name, nil
	}
}

func (d *Handle) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}

func (h *Handle) GetRootValues() ([]*Value, error) {
	values, err := h.col.GetRootValues(h.setValue)
	return values, err
}

func (h *Handle) setColbyValue(r *reflect.Value) {
	if h.col != nil {
		return
	}
	h.col = h.db.GetColByName(r.Type().Name())
}

func (h *Handle) addSetValue(s *setValue) {
	h.setValueLst = append(h.setValueLst, s)
}

func (h *Handle) getInsertValue() *reflect.Value {
	return h.setValueLst[0].value
}

func (h *Handle) setResult(i interface{}) {
	h.result = *newResult(i)
}
