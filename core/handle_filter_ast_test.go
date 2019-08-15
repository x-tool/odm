package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_AST(t *testing.T) {
	tree, err := setBracketsTree("a = b")
	str, err := json.Marshal(tree)
	fmt.Println(string(str), err)
}


odm.Where(
	odm.And(
		odm.And(value, "in", a),
		odm.And()
	)
)

odm.link(odm.link{
	
}
)