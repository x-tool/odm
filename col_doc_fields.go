package odm

type DocFieldLst []*DocField

func (d *DocFieldLst) getFieldByName(name string) (o *DocFields) {
	for _, v := range d {
		if v.Name == name {
			o = append(o, v)
		}
	}
	return
}

func (d *DocFieldLst) getRootFieldLst() (rd *DocFieldLst){
	for _,v := range d {
		if v.pid = rootPid {
			rd = append(rd, v)
		}
	}
	return 
}

func (d *DocFieldLst) getExtendFieldLst() (rd *DocFieldLst){
	for _,v := range d {
		if v.isExtend() {
			rd = append(rd, v)
		}
	}
	return 
}