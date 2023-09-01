package util

import "strconv"

// ToInt Convert string to int type
func ToInt(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		return -1
	}

	return n
}

func ToInt64(v string) int64 {
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return -1
	}

	return n
}
