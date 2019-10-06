package db

import (
	"strings"
)

func InStringList(field string, s []string) string {
	if len(s) == 0 {
		return "1=0"
	}

	escaped := make([]string, len(s))
	for i, v := range s {
		escaped[i] = Escape(v)
	}
	return field + " IN (" + strings.Join(escaped, ",") + ")"
}
