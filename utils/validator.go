package utils

import "strings"

func IsValidPriority(p string) bool {
	switch strings.ToLower(p) {
	case "low", "medium", "high":
		return true
	}
	return false
}

func IsValidStatus(status string) bool {
	switch strings.ToLower(status) {
	case "new", "progress", "completed":
		return true
	}
	return false
}
