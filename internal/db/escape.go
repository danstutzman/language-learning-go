package db

import (
	"gopkg.in/guregu/null.v3"
	"strconv"
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
	return "'" + time.UTC().Format("2006-01-02T15:04:05Z") + "'"
}

func EscapeNullTime(time null.Time) string {
	if !time.Valid {
		return "NULL"
	}

	return EscapeTime(time.Time)
}

func EscapeBool(b bool) string {
	if b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func EscapeNullBool(b null.Bool) string {
	if !b.Valid {
		return "NULL"
	}

	return EscapeBool(b.Bool)
}

func EscapeNullInt(i null.Int) string {
	if !i.Valid {
		return "NULL"
	}

	return strconv.FormatInt(i.Int64, 10)
}
