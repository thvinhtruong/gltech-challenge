package conversion

import "strings"

func MergeString(strs ...string) string {
	return strings.Join(strs, ",")
}
