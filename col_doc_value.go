package odm

import "reflect"

func (d *Doc) newItem() (v reflect.Value) {
	return reflect.New(*(d.t))
}
