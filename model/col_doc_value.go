package model

import "reflect"

func (d *Doc) newItemValue() (v *reflect.Value) {
	_v := reflect.New(*(d.t))
	return &_v
}

func (d *Doc) getNewItemRootValue(rootV *reflect.Value) (returnVLst []*reflect.Value) {
	docFields := d.getRootDetails()
	for _, v := range docFields {
		returnVLst = append(returnVLst, d.getRootDetailValue(rootV, v))
	}
	return
}
