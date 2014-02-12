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

func ItoB5(bits uint64) [5]byte {
	return [5]byte{
		byte(bits >> 32),
		byte(bits >> 24),
		byte(bits >> 16),
		byte(bits >> 8),
		byte(bits),
	}
}

func B5(bits string) [5]byte {
	i, _ := strconv.ParseUint(bits, 2, 64)

	return ItoB5(i)
}
