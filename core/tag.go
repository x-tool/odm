package core

import (
	"strings"
)

// this vars could be user modify, so use var not const
var (
	tagName          = "xodm"
	tagItemSeparator = ":" // Ex: xodm:"name:value"
	tagSeparator     = ";" // Ex: xodm:"name1:value1;name2:value2"
	tagMark          = "@" // Ex: xodm:"@mark"
)

// `xodm:"@sign"`
type odmTag struct {
	sourceTag string
	mark      string // find docfield quick by custom string
	allowNull bool
	lst       map[string]string
}

func newTag(s string) *odmTag {
	_o := &odmTag{}
	_o.lst = make(map[string]string)
	_s := strings.TrimSpace(s)
	_o.sourceTag = _s
	_o.allowNull = true
	lst := strings.Split(_s, tagSeparator)
	for _, v := range lst {
		fieldLst := strings.Split(v, tagItemSeparator)
		fieldLstLen := len(fieldLst)
		var name string
		var value string
		// tagMark value like @xxx, so format it diffrent with other
		if fieldLstLen == 1 {
			var _str = fieldLst[0]
			switch {
			// notnull
			case _str == "notnull":
				_o.allowNull = false
				name = "notnull"
				value = ""
				// @xxx
			case strings.Index(_str, tagMark) == 0:
				_o.mark = string([]rune(fieldLst[0])[1:])
				name = tagMark
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
