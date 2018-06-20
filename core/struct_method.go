package core

import "strings"

func (d *odmStruct) getFieldByName(name string) (o structFieldLst) {
	return d.fieldNameMap[name]
}

func (d *odmStruct) getFieldByTag(tag string) (o *structField) {
	return d.fieldTagMap[tag]
}

// "path.fieldName"
func (d *odmStruct) getFieldByPath(pathStr string) (f *structField) {
	// check dependLst
	path := strings.Split(pathStr, ".")
	fieldNamme := path[len(path)-1]
	dependLst := path[len(path)-1:]
	fields := d.getFieldByName(fieldNamme)

	if len(fields) == 1 {
		f = fields[0]
	} else {
		f = d.getFieldByDependLst(dependLst, fieldNamme)
	}
	return
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

func (d *odmStruct) getFieldByDependLst(Lst []string, fieldName string) (_field *structField) {
	// fields := d.getFieldByName(fieldName)
	for _, field := range d.fields {
		var check bool = false
		for i, dependField := range field.dependLst {
			if dependField.Name() != Lst[i] {
				check = false
				break
			} else {
				check = true
			}
		}
		if check == true {
			_field = field
			break
		}
	}
	return
}
