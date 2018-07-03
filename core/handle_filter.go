package core

// type filters []filter

// type handleType int

// const (
// 	insertData handleType = iota
// 	updateData
// 	deleteData
// 	queryData
// )

// type filterCompare string

// const (
// 	sameCompare      filterCompare = "like"  // like
// 	sameCompareLeft                = "?like" // ??like
// 	sameCompareRight               = "like?" // like??
// 	equalCompare                   = "=="    // ==
// 	isNullCompare                  = "isNull"
// 	betweenCompare                 = "between"
// 	inCompare                      = "in"
// )

// type filterJoin string

// const (
// 	andFilter filterJoin = "and"
// 	orFilter             = "or"
// 	notFilter            = "not"
// )

// type filterItem struct {
// 	target       *structField
// 	compare      filterCompare
// 	value        interface{}
// 	connect      filterJoin
// 	childFilters filters
// }

// type filter struct {
// 	Handle    *Handle
// 	queryKind string
// 	queryV    *reflect.Value
// 	modeV     *reflect.Value
// 	queryLst  []filterItem
// 	limitNum  int
// 	limitDesc bool
// }

// func newFilter(o *Handle) *filter {
// 	r := &filter{
// 		Handle: o,
// 		// filterV:    rV,
// 		// filterKind: t,
// 	}
// 	// r.setDependToDoc()
// 	return r
// }

// func newfilterWithoutCol(rV *reflect.Value) *filter {
// 	r := &filter{
// 		filterV: rV,
// 	}
// 	return r
// }

// func (r *filter) setDependToDoc() {
// 	// T := r.filterV.Type()
// 	// var value reflect.Type
// 	// if T.Kind() == reflect.Slice {
// 	// 	value = T.Elem()
// 	// } else {
// 	// 	value = T
// 	// }
// 	// var valueItem reflect.Value
// 	// var valueItemT reflect.Type
// 	// if value.Kind() == reflect.Slice {
// 	// 	valueItem = r.filterV.Elem()
// 	// } else {
// 	// 	valueItem = *r.filterV
// 	// }
// 	// valueItemT = valueItem.Type()
// 	// for i := 0; i < valueItem.NumField(); i++ {
// 	// 	field := valueItem.Field(i)
// 	// 	fieldT := valueItemT.Field(i)
// 	// 	if isDocMode(fieldT.Name) {
// 	// 		r.modeV = &field
// 	// 	}
// 	// 	newfilterItem := r.Handle.DependToDoc(strings.Split(fieldT.Tag.Get(tagName), "."), fieldT.Name)
// 	// 	r.filterLst = append(r.filterLst, newfilterItem)
// 	// }

// }
