package typeman

type Config struct {
	IsValid func(i interface{}) bool
}
type TypeManager struct {
	typeLst []Type
}

func (t *TypeManager) Register(typ interface{}, conf Config) (err error) {

	return
}
