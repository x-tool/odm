package typeman

type TypeStruct struct {
	Name string
	Kind
	MaxLen   int64
	MinLen   int64
	isFixLen bool
	// use for group type
	keyType   *TypeStruct
	valueType *TypeStruct
	fieldLst  *TypeStruct
}
