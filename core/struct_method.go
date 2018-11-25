package core

func (d *odmStruct) getFieldByName(name string) (o structFieldLst) {
	return d.fieldNameMap[name]
}

func (d *odmStruct) getFieldByMark(tag string) (o *structField) {
	return d.fieldMarkMap[tag]
}

func (d *odmStruct) GetRootFields() structFieldLst {
	return d.rootFields
}

func (d *odmStruct) getExtendFields() (lst structFieldLst) {
	for _, v := range d.fields {
		if v.isExtend {
			lst = append(lst, v)
		}
	}
	return
}
