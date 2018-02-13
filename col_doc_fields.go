package odm

type docFieldLst []*DocField

func (d *DocFieldLst) getFieldsByName(name string) (o *docFields) {
	for _, v := range d {
		if v.Name == name {
			o = append(o, v)
		}
	}
	return
}

func (d *docFieldLst) getRootFieldLst() (rd *docFieldLst){
	for _,v := range d {
		if v.pid = rootPid {
			rd = append(rd, v)
		}
	}
	return 
}

func (d *docFieldLst) getExtendFieldLst() (rd *docFieldLst){
	for _,v := range d {
		if v.IsExtend() {
			rd = append(rd, v)
		}
	}
	return 
}

func (d *docFieldLst) getSingleTypeFieldLst() (rd *docFieldLst){
	for _,v := range d {
		if v.isSingleType() {
			rd = append(rd, v)
		}
	}
	return 
}
func (d *docFieldLst) getGroupTypeFieldLst() (rd *docFieldLst){
	for _,v := range d {
		if v.isGroupType() {
			rd = append(rd, v)
		}
	}
	return 
}