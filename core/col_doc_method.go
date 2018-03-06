package core

func (d *doc) GetRootFields() (lst docFieldLst) {
	for _, v := range d.fields {
		if v.extendParent == nil && v.IsExtend() == false {
			lst = append(lst, v)
		}
	}
	return
}
