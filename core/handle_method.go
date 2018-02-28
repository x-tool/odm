package core

import (
	"errors"
	"strings"

	"github.com/x-tool/tool"
)

func (r *Handle) dependtoDocOneStr(s string) (d *docField) {
	_s := strings.Split(s, ".")
	if len(_s) > 1 {
		return r.DependToDoc(_s[:len(_s)-2], _s[len(_s)-1])
	} else {
		return r.DependToDoc([]string{}, _s[0])
	}
}
func (r *Handle) DependToDoc(dependLst []string, name string) (d *docField) {
	if len(dependLst) == 0 {
		field := r.Col.doc.getFieldByName(name)
		if len(field) != 1 {
			tool.Panic("Handle", errors.New("name not be single, you should add dependLst to find doc field"))
		} else {
			return field[0]
		}
	} else {
		docFieldLst := r.Col.Doc.getFieldByName(name)
		for _, val := range docFieldLst {
			if len(dependLst) != len(val.dependLst) {
				continue
			}
			var check bool = true
			for i, _ := range val.dependLst {
				if val.dependLst[i].Name != dependLst[i] {
					check = false
					break
				}
			}
			if check {
				d = val
				break
			}
		}
	}
	return
}
