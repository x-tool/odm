package core

import (
	"strings"
)

var (
	tagName      = "xodm"
	tagSeparator = ";"
	tag_Tag     = "@"
)

// `xodm:"@sign"`
type odmTag struct {
	sourceTag string
	sign      string // find docfield quick by custom string
}

func newTag(s string) *odmTag {
	_o := &odmTag{}
	_s := strings.TrimSpace(s)
	_o.sourceTag = _s
	lst := strings.Split(_s, tagSeparator)
	for _, v := range lst {
		fieldLst := strings.Split(v, ":")
		fieldLstLen := len(fieldLst)
		if fieldLstLen == 1 {
			if strings.Index(fieldLst[0], tag_Tag) == 0 {
				_o.sign = string([]rune(fieldLst[0])[1:])
			}
		} else if fieldLstLen == 2 {
			name := strings.TrimSpace(fieldLst[0])
			value := strings.TrimSpace(fieldLst[1])
			switch name {
			case tag_Tag:
				_o.sign = value
			}
		}
	}
	return _o
}
