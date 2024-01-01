package utils

import (
	"strconv"
)

func StringToInt(s string) int64 {
	integer, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Error during conversion")
	}

	return integer
}
