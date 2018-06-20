package core

import (
	"strings"
)

var (
	tagName = "xodm"
	tag_Ptr = "p"
)

type odmTag struct {
	sourceTag string
	sign      string // find docfield quick by custom string
}

func newTag(s string) *odmTag {
	_o := &odmTag{}
	_s := strings.TrimSpace(s)
	_o.sourceTag = _s
	lst := strings.Split(_s, " ")
	for _, v := range lst {
		fieldLst := strings.Split(v, ":")
		switch fieldLst[0] {
		case tag_Ptr:
			_o.sign = fieldLst[1]
		}
	}
	return _o
}
