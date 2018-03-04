package core

func (d *doc) getRootDetails() (lst dependLst) {
	for _, v := range d.fields {
		if v.extendParent == nil {
			lst = append(lst, v)
		}
	}
	return
}
