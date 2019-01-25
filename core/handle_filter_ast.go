package core

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	bracketLeft        = "("
	bracketRight       = ")"
	filterNot          = "not"
	filterAnd          = "and"
	filterOr           = "or"
	placeholderValue   = "?"
	oprationEQ         = "="
	oprationNEQ        = "<>"
	oprationNEQ2       = "!="
	oprationGT         = ">"
	oprationEGT        = ">="
	oprationLT         = "<"
	oprationELT        = "<="
	oprationLike       = "like"
	oprationBETWEEN    = "between"
	oprationBETWEENAnd = "and"
	oprationIN         = "in"
	oprationNOTNULL    = "is not null"
	oprationNULL       = "is null"
)

const (
	codeNot = "!"
	codeAnd = "&&"
	codeOr  = "||"
	codeEQ  = "=="
)

var defaultVarsRegexp *regexp.Regexp

type CompareKind int

const (
	CompareAnd = iota
	CompareOr
	CompareNot
)

type ASTTree struct {
	parent *ASTTree
	source string
	Link   string
	child  []*ASTTree
	Field  *StructField
	CompareKind
	valueLst []interface{}
}

func (a *ASTTree) IsBox() bool {
	return len(a.child) != 0
}

func setBracketsTree(s string, values ...interface{}) (rootTree *ASTTree, err error) {
	rootTree.source = s
	focusTree := rootTree
	var valuesIndex = 0
	var valuesLen = len(values)
	var getValue = func() (interface{}, error) {
		if valuesIndex+1 >= valuesLen {
			return nil, errors.New("input values lenght less than '?' in string")
		}
		return values, nil
	}
	var state_string bool
	var state_string_esc bool
	var state_number bool
	var state_func bool
	var state_field bool
	var valueStartIndex int
	for i, v := range s {
		letter := string(v)
		if state_string {
			if state_string_esc {
				state_string_esc = false
				continue
			}
			if letter == "\"" {
				state_string = false
				focusTree.valueLst = append(focusTree.valueLst, s[valueStartIndex+1:i])
			}
			continue
		}
		if state_field {
			switch letter {
			case " ":
				state_field = false
			case "(":
				state_field = false
				state_func = true
			}
			continue
		}
		if state_func {
			switch letter {
			case ")":
				state_func = false
			}
			continue
		}
		if state_number {
			switch letter {
			case " ":
				state_number = false
			}
			continue
		}
		findStrIndexs := defaultVarsRegexp.FindStringSubmatchIndex(s[i:])
		if findStrIndexs != nil {
			_v := focusTree.source[findStrIndexs[0]:findStrIndexs[1]]
			switch _v {
			case bracketLeft:
				newTree := &ASTTree{source: focusTree.source[findStrIndexs[1]:]}
				focusTree.child = append(focusTree.child, newTree)
				focusTree = newTree
			case bracketRight:
				focusTree.source = focusTree.source[:findStrIndexs[0]]
				// focusTree.parent.source =
			case placeholderValue:
				value, err := getValue()
				if err != nil {
					return nil, err
				}
				focusTree.valueLst = append(focusTree.valueLst, value)
			case filterNot:
				fallthrough
			case codeNot:
				focusTree.CompareKind = CompareNot
			case filterAnd:
				fallthrough
			case codeAnd:
				focusTree.CompareKind = CompareAnd
			case filterOr:
				fallthrough
			case codeOr:
				focusTree.CompareKind = CompareOr
			case codeEQ:
				_v = oprationEQ
				fallthrough
			case oprationEQ:
				fallthrough
			case oprationNEQ:
				fallthrough
			case oprationNEQ2:
				fallthrough
			case oprationGT:
				fallthrough
			case oprationEGT:
				fallthrough
			case oprationLT:
				fallthrough
			case oprationELT:
				fallthrough
			case oprationNOTNULL:
				fallthrough
			case oprationNULL:
				fallthrough
			case oprationIN:
				fallthrough
			case oprationBETWEEN:
				fallthrough
			case oprationLike:
				focusTree.Link = _v
			}

		} else {
			switch letter {
			case "\"":
				state_string = true
			case "0":
				fallthrough
			case "1":
				fallthrough
			case "2":
				fallthrough
			case "3":
				fallthrough
			case "4":
				fallthrough
			case "5":
				fallthrough
			case "6":
				fallthrough
			case "7":
				fallthrough
			case "8":
				fallthrough
			case "9":
				fallthrough
			case ".":
				state_number = true
			default:
				state_field = true
			}
			continue
		}
	}
	return
}

func init() {
	// init defaultVarsRegexp
	defaultVarsRegexpLst := []string{
		bracketLeft,
		bracketRight,
		filterNot,
		filterAnd,
		filterOr,
		placeholderValue,
		oprationEQ,
		oprationNEQ,
		oprationNEQ2,
		oprationGT,
		oprationEGT,
		oprationLT,
		oprationELT,
		oprationLike,
		oprationBETWEEN,
		oprationNOTNULL,
		oprationNULL,
		codeNot,
		codeAnd,
		codeOr,
		codeEQ,
	}
	var rangeDefaultStr string
	for i, v := range defaultVarsRegexpLst {
		if i == 0 {
			rangeDefaultStr = v
		} else {
			rangeDefaultStr = rangeDefaultStr + "|" + v
		}
	}
	defaultVarsRegexp, _ = regexp.Compile(fmt.Sprintf(" *^(%v)", rangeDefaultStr))
}
