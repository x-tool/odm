package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_AST(t *testing.T) {
	tree, err := setBracketsTree("a = b")

	fmt.Println(json.Marshal(tree), err)
}
