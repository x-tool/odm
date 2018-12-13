package core

import (
	"strings"
)

// this vars could be user modify, so use var not const
var (
	tagName          = "xodm"
	tagItemSeparator = ":"      // Ex: xodm:"name:value"
	tagSeparator     = ";"      // Ex: xodm:"name1:value1;name2:value2"
	tagMark          = "@"      // Ex: xodm:"@mark"
	tagPath          = "__path" // Ex: xodm:"path" default name to odmTag lst map key
)

// `xodm:"@sign"`
type odmTag struct {
	sourceTag string
	mark      string // find docfield quick by custom string
	allowNull bool
	path      string // odmPath
	lst       map[string]string
}

func newTag(s string) *odmTag {
	_o := &odmTag{}
	_o.lst = make(map[string]string)
	_s := strings.TrimSpace(s)
	_o.sourceTag = _s
	lst := strings.Split(_s, tagSeparator)
	for _, v := range lst {
		fieldLst := strings.Split(v, tagItemSeparator)
		fieldLstLen := len(fieldLst)
		var name string
		var value string
		// tagMark value like @xxx, so format it diffrent with other
		if fieldLstLen == 1 {
			if strings.Index(fieldLst[0], tagMark) == 0 {
				_o.mark = string([]rune(fieldLst[0])[1:])
				name = tagMark
				value = _o.mark
			} else {
				_o.path = fieldLst[0]
				name = tagPath
				value = _o.mark
			}
		} else if fieldLstLen == 2 {
			name := strings.TrimSpace(fieldLst[0])
			value := strings.TrimSpace(fieldLst[1])
			switch name {
			case tagMark:
				_o.mark = value
			}
		}
		_o.lst[name] = value
	}
	return _o
}

func (t *odmTag) AllowNull() bool {
	return t.allowNull
}
