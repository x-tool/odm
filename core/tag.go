package core

import (
	"strings"
)

// this vars could be user modify, so use var not const
var (
	tagName      = "xodm"
	tagSeparator = ";"
	tagMark      = "@"
)

// `xodm:"@sign"`
type odmTag struct {
	sourceTag string
	mark      string // find docfield quick by custom string
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
			if strings.Index(fieldLst[0], tagMark) == 0 {
				_o.mark = string([]rune(fieldLst[0])[1:])
			}
		} else if fieldLstLen == 2 {
			name := strings.TrimSpace(fieldLst[0])
			value := strings.TrimSpace(fieldLst[1])
			switch name {
			case tagMark:
				_o.mark = value
			}
		}
	}
	return _o
}
