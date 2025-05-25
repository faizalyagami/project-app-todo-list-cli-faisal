package utils

import "strings"

func FormatTitle(title string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(title)))
}
