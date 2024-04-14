package util

import (
	"strings"
)

func ToSnakeCase(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "_")
	return str
}
