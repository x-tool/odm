package core

import (
	"fmt"
	"regexp"
)

func main() {
	stringExp, _ := regexp.Compile(`'.*[^\'|.].*'|".*[^\"|.].*"`)
	stringIndexLst := stringExp.FindAllStringIndex(`aa@ff|ss + bb = "cc rr ss\'kk\'\'\""`, -1)

	odmFiledExp, _ := regexp.Compile(`\b[\w|@|\||.]+`)
	fieldIndexLst := odmFiledExp.FindAllString(`aa@ff|ss.bb + bb = "cc rr ss\'kk\'\'\""`, -1)
	fmt.Println(stringIndexLst)
	fmt.Println(fieldIndexLst)
	fmt.Println("ss")
}
