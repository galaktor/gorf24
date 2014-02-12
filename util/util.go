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

func B4(bits string) uint32 {
	i, _ := strconv.ParseUint(bits, 2, 32)
	return uint32(i)
}

func B5(bits string) [5]byte {
	i, _ := strconv.ParseUint(bits, 2, 64)

	return [5]byte {
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}










