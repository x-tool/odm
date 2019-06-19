package core

const (
	bracketLeft        = "("
	bracketRight       = ")"
	filterNot          = "not"
	filterAnd          = "and"
	filterOr           = "or"
	placeholderValue   = "?"
	oprationEQ         = "="
	oprationNEQ        = "!="
	oprationGT         = ">"
	oprationEGT        = ">="
	oprationLT         = "<"
	oprationELT        = "<="
	oprationLike       = "like"
	oprationBETWEEN    = "between"
	oprationBETWEENAnd = "and"
	oprationIN         = "in"
	oprationNOTNULL    = "notNull"
	oprationNULL       = "null"
)

const (
	codeNot = "!"
	codeAnd = "&&"
	codeOr  = "||"
	codeEQ  = "=="
)

type CompareKind int

const (
	CompareAnd = iota
	CompareOr
	CompareNot
)

type ASTTreeNode struct {
	parent      *ASTTreeNode
	source      string
	Link        string
	leftValue   interface{}
	RightValues []interface{}
	child       []*ASTTreeNode
	CompareKind
}

func createASTtree(lst ...interface{}) (tree *ASTTreeNode, err error) {
	tree = &ASTTreeNode{}
	for _, v := range lst {
		_, err = tree.addChild(v)
		if err != nil {
			return
		}
	}
	return
}
