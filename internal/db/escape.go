package db

import (
	"strings"
)

func Escape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}
