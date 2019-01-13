package core

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
	str := strings.TrimSpace(box.source)
	var checkBrackets bool
	lexer()
	reg, err := regexp.Compile(bracketLeft + "|" + bracketRight)
	return
}

type lexerKind int

const (
	bracketLeft lexerKind = iota
	bracketRight
	lexerField
	link
)

type lexerBox struct {
	
}

func lexer(s string) ()
func getEndBracketIndex(s string, index) (i int, err error) {
	var l = len(s[index:])
	for i,v:= range l{

	}
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
