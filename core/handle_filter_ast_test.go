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
	db.Link(
		"a", "b", 
		odm.Where(
			odm.And("o.bb in ?", a),
			odm.And()
		)
	).alias("z")

)

odm.link(odm.link{
	
}
)