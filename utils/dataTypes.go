package utils

import "strconv"

// GetInt parses a string to integer
func GetInt(val string) int {
	res, _ := strconv.Atoi(val)
	return res
}
