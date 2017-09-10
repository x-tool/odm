package xodm

import "regexp"

func tagExtend(tag string) bool {
	matched, err := regexp.MatchString("extend", tag)
	if err != nil {
		return false
	}
	return matched
}
