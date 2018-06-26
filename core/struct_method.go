package core

import "strings"

func (d *odmStruct) getFieldByName(name string) (o structFieldLst) {
	return d.fieldNameMap[name]
}

func (d *odmStruct) getFieldByTag(tag string) (o *structField) {
	return d.fieldTagMap[tag]
}

// "fieldName"
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
		f = d.getFieldByDependLst(fieldNamme, dependLst)
	}
	return
}

func (d *odmStruct) getFieldByDependLst(fieldName string, Lst []string) (_field *structField) {
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

// "@tag"
// "fieldName"
// "path.fieldName"
func (d *odmStruct) getFieldByString(str string)(f *structField, err error) {
	var sign = str[:1]
	if sign == tag_Tag {
		f = d.getFieldByTag(str[1:])
	} else {
		f = d.getFieldByPath(str)
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

