package gorf24

import (
	"testing"

	"."
)

func TestByte_Zero_ReturnsZero(t *testing.T) {
	c := gorf24.Command(gorf24.B("00000000"))
	expected := gorf24.B("00000000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_AllOnes_ReturnsSame(t *testing.T) {
	c := gorf24.Command(gorf24.B("11111111"))
	expected := gorf24.B("11111111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_LSByteZero_ReturnsSame(t *testing.T) {
	c := gorf24.Command(gorf24.B("11110000"))
	expected := gorf24.B("11110000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_MSByteZero_ReturnsSame(t *testing.T) {
	c := gorf24.Command(gorf24.B("00001111"))
	expected := gorf24.B("00001111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}



















