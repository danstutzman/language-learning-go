package db

import (
	"strings"
)

func Escape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

func EscapePtr(s *string) string {
	if s == nil {
		return "NULL"
	} else {
		return "'" + strings.ReplaceAll(*s, "'", "''") + "'"
	}
}
