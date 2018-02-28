package core

type docFieldLst []*docField
type dependLst docFieldLst

func (d *docFieldLst) getFieldsByName(name string) (o docFieldLst) {
	for _, v := range *d {
		if v.GetName() == name {
			o = append(o, v)
		}
	}
	return
}

func (d *docFieldLst) getRootFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.pid == rootPid {
			rd = append(rd, v)
		}
	}
	return
}

func (d *docFieldLst) getExtendFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.IsExtend() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *docFieldLst) getSingleTypeFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.isSingleType() {
			rd = append(rd, v)
		}
	}
	return
}
func (d *docFieldLst) getGroupTypeFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.isGroupType() {
			rd = append(rd, v)
		}
	}
	return
}
