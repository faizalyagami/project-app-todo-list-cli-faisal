package utils

import "strings"

func IsValidPriority(p string) bool {
	switch strings.ToLower(p) {
	case "low", "medium", "high":
		return true
	}
	return false
}
