package core

import (
	"fmt"
	"reflect"
)

const (
	bracketLeft  = "("
	bracketRight = ")"
	codeNot      = "!"
	codeAnd      = "&&"
	codeOr       = "||"
	filterNot    = "not"
	filterAnd    = "and"
	filterOr     = "or"
)

type CompareKind int

const (
	likeCompare    CompareKind = iota
	equalCompare               // ==
	isNullCompare              // isNull
	betweenCompare             // between
	inCompare                  // in
)

type linkKind int

type ASTTree struct {
	source string // temp user iput
	linkKind
	child []ASTTree
	field *StructField
	CompareKind
	valueLst []reflect.Value
}

func setBracketsTree(s string) (err error) {
	lexerBoxs, err := lexerAnalysis(s)
	if err != nil {
		return err
	}

	return
}

type lexerKind int

const (
	lexerBracketLeft lexerKind = iota
	lexerBracketRight
	lexerField
	lexerAnd
	lexerOr
	lexerNot
)

type lexerBox struct {
}

func lexerAnalysis(s string) (boxs []lexerBox, err error) {
	var state_string bool
	var state_string_esc bool
	var state_func bool
	for i, v := range s {
		letter := string(v)
		// in out state_string_esc
		if !state_string_esc && letter == "\\" {
			state_string_esc = true
			continue
		}
		if state_string_esc {
			state_string_esc = false
			continue
		}
		// in out state_string
		if letter == "\"" {
			if !state_string {
				state_string = true
				continue
			}
			if state_string && !state_string_esc {
				state_string = false
				continue
			}
		}

	}
	return
}

func init() {
	// init formatBracketStr
	formatBracketStrLst := []string{
		codeNot,
		codeAnd,
		codeOr,
		filterNot,
		filterAnd,
		filterOr,
	}
	_regStr := fmt.Sprintf(" *%v|^ *%v", bracketLeft, bracketLeft)
	var beforeBracketLeft string
	for i, v := range formatBracketStrLst {
		if i == 0 {
			beforeBracketLeft = v
		} else {
			beforeBracketLeft = beforeBracketLeft + "|" + v
		}
	}
	formatBracketStr = fmt.Sprintf("[%v]%v", beforeBracketLeft, _regStr)
}
