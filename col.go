package xodm

type ColLst []Col

type ColInterface interface {
	Name() string
}

type Col struct {
	Name      string
	Tree      Coltree
	detailLst []*ColDetailAttribute
}

type ColDetailAttribute struct {
	Name       string
	DetailType string
	ParentProp []string
}

type Coltree struct {
	Ttype  string
	DBtype string
	sign   string
}
