package db

import (
	"database/sql"
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
func EscapeNullString(s sql.NullString) string {
	if !s.Valid {
		return "NULL"
	} else {
		return "'" + strings.ReplaceAll(s.String, "'", "''") + "'"
	}
}
