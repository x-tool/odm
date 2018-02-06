package odm

import "reflect"

type docField struct {
	name      string
	selfType  string
	dbType    string
	id        int
	pid       int // field golang parent real ID; default:-1
	isExtend  bool
	extendPid int // field Handle parent ID; default:-1
	dependLst
	Tag     string
	funcLst map[string]string
}

func (d *docField) IsExtend() bool {
	return d.isExtend
}

func (o *docField) getRootFieldDB() (r *docField) {
	switch len(o.dependLst) {
	case 0:
		return 0
	default:
		if o.dependLst[0].isExtend {
			return o.dependLst[1]
		} else {
			return o.dependLst[0]
		}
	}
}

func (o *docField) getDependLstDB() (r docFieldLst) {
	for _, v := range o.dependLst {
		if v.isExtend {
			r = append(r, v)
		}
	}
	return
}

func newdocField(d *doc, t *reflect.StructField, Pid int, extendPid int) {
	fieldType := *t
	fieldTypeStr := formatTypeToString(&fieldType.Type)
	id := len(d.fields)
	tag := fieldType.Tag.Get(tagName)
	isExtend := checkdocFieldisExtend(tag)
	extendField := d.getFieldById(extendPid)
	var dependLst dependLst
	if extendField == nil {
	} else {
		dependLst = append(extendField.dependLst, extendField)
	}

	field := &docField{
		Name:      fieldType.Name,
		Type:      fieldTypeStr,
		DBType:    d.col.DB.SwitchType(fieldTypeStr),
		Id:        id,
		Pid:       Pid,
		isExtend:  isExtend,
		dependLst: dependLst,
		extendPid: extendPid,
		Tag:       tag,
	}
	d.fields = append(d.fields, field)
	switch t.Type.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		_fieldType := fieldType.Type.Elem()
		count := _fieldType.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = Pid
			} else {
				extendPid = id
			}
			field := _fieldType.Field(i)
			newdocField(d, &field, id, extendPid)
		}
	case reflect.Struct:
		// if time package not range time struct
		if t.Type.PkgPath() == "time" {
			return
		}
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = Pid
			} else {
				extendPid = id
			}
			field := fieldType.Type.Field(i)
			newdocField(d, &field, id, extendPid)
		}

	}
}
