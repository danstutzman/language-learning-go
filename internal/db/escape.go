package db

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
	"strings"
	"time"
)

func Escape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

func EscapePtr(s *string) string {
	if s == nil {
		return "NULL"
	}

	return Escape(*s)
}

func EscapeNullString(s sql.NullString) string {
	if !s.Valid {
		return "NULL"
	}

	return Escape(s.String)
}

func EscapeTime(time time.Time) string {
	return "'" + time.Format("2006-01-02T15:04:05Z") + "'"
}

func EscapeNullTime(time null.Time) string {
	if !time.Valid {
		return "NULL"
	}

	return EscapeTime(time.Time)
}
