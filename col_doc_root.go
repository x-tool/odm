package odm

func (doc *Doc) getRootExtendFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.Pid == -1 && v.extendPid != -1 && v.isExtend {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootSinpleFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid == -1 && !v.isExtend && !doc.checkComplexField(v) {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootComplexFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid == -1 && !v.isExtend && doc.checkComplexField(v) {
			d = append(d, v)
		}
	}
	return
}

func (d *Doc) getRootDetails() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !v.isExtend {
			doc = append(doc, v)
		}
	}
	return
}
func (d *Doc) getRootDetailsWithoutExtend() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 {
			doc = append(doc, v)
		}
	}
	return
}
