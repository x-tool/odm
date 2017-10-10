package odm

import "regexp"
import "reflect"

func tagIsExtend(tag string) bool {
	matched, err := regexp.MatchString("extend", tag)
	if err != nil {
		return false
	}
	return matched
}

func tagisExtendField(r *reflect.StructField) bool {
	tag := r.Tag.Get(tagName)
	return tagIsExtend(tag)
}
