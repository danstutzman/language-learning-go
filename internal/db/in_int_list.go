package db

import (
	"strconv"
	"strings"
)

func InIntList(field string, i []int) string {
	if len(i) == 0 {
		return "1=0"
	}

	s := make([]string, len(i))
	for i, v := range i {
		s[i] = strconv.Itoa(v)
	}
	return field + " IN (" + strings.Join(s, ",") + ")"
}
