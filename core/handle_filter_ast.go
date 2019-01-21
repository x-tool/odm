package core

import (
	"fmt"
	"reflect"
	"regexp"
)

const (
	bracketLeft     = "("
	bracketRight    = ")"
	filterNot       = "not"
	filterAnd       = "and"
	filterOr        = "or"
	stringValue     = "?"
	oprationEQ      = "="
	oprationNEQ     = "<>"
	oprationNEQ2    = "!="
	oprationGT      = ">"
	oprationEGT     = ">="
	oprationLT      = "<"
	oprationELT     = "<="
	oprationLike    = "like"
	oprationBETWEEN = "between"
	oprationNOTNULL = "is not null"
	oprationNULL    = "is null"
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
	source string // temp user iput
	link   string
	child  []ASTTree
	field  *StructField
	CompareKind
	valueLst []reflect.Value
}

func setBracketsTree(s string) (rootTree ASTTree, err error) {
	var state_string bool
	var state_string_esc bool
	var state_number bool
	var state_func bool
	var state_field bool
	for i, v := range s {
		letter := string(v)
		if state_string {
			if state_string_esc {
				state_string_esc = false
				continue
			}
			if letter == "\"" {
				state_string = false
			}
			continue
		}
		if state_field  {
			switch letter{
			case " ":
				state_field = false
			case "(" :
				state_field = false
				state_func = true
			}
			continue
		}
		if state_func {
			switch letter{
			case ")" :
				state_func = false
			}
			continue
		}
		if state_number {
			switch letter{
				case " "
				state_number = false
			}
			continue
		}
		findStr := defaultVarsRegexp.FindStringSubmatch(s[i:])
		if findStr != nil {

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
		stringValue,
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
