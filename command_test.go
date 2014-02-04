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

func TestByte_Zero_ReturnsZero(t *testing.T) {
	c := gorf24.Command(b("00000000"))
	expected := b("00000000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_AllOnes_ReturnsSame(t *testing.T) {
	c := gorf24.Command(b("11111111"))
	expected := b("11111111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_LSByteZero_ReturnsSame(t *testing.T) {
	c := gorf24.Command(b("11110000"))
	expected := b("11110000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_MSByteZero_ReturnsSame(t *testing.T) {
	c := gorf24.Command(b("00001111"))
	expected := b("00001111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}



















