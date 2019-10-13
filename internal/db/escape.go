package db

import (
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

func EscapeNullString(s null.String) string {
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
