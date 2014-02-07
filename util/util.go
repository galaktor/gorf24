package util

import (
	"strconv"
)

/*
 parse a string representation of bits into
 a byte; for easier testing
*/
func B(bits string) byte {
	i, _ := strconv.ParseUint(bits, 2, 8)
	return byte(i)
}