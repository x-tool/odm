package core

func (d *doc) GetRootDetails() (lst dependLst) {
	for _, v := range d.fields {
		if v.extendParent == nil {
			lst = append(lst, v)
		}
	}
	return
}
