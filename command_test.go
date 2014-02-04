package gorf24

import (
	"testing"
	"strconv"

	"."
)

/* 
 parse a string representation of bits into
 a byte; for easier testing
*/
func b(bits string) byte {
	i,_ := strconv.ParseUint(bits, 2, 8)
	return byte(i)
}

