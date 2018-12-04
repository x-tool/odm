package core

import "reflect"

func (d *odmStruct) getFieldByName(name string) (o structFieldLst) {
	return d.fieldNameMap[name]
}

func (d *odmStruct) getFieldByMark(tag string) (o *structField) {
	return d.fieldMarkMap[tag]
}

func (d *odmStruct) getExtendFields() (lst structFieldLst) {
	for _, v := range d.fields {
		if v.isExtend {
			lst = append(lst, v)
		}
	}
	return
}

func (d *odmStruct) GetRootFields() structFieldLst {
	return d.rootFields
}

/// ***** should optimize, if complex nesting
func (d *odmStruct) GetRootValues(rootValue *reflect.Value) (result []reflect.Value) {
	for _, v := range d.rootFields {
		var value = *rootValue
		for _, _v := range v.dependLst {
			value = value.FieldByName(_v.name)
		}
		value = value.FieldByName(v.name)
		result = append(result, value)
	}
	return
}
