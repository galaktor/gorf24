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

func TestReadCmd_ZeroBits_ReturnsZero(t *testing.T) {
	r := gorf24.Register(b("00000000"))
	expected := b("00000000")

	result := r.ReadCmd().Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestReadCmd_AllOnes_ReturnsLastFiveBits(t *testing.T) {
	r := gorf24.Register(b("11111111"))
	expected := b("00011111")

	result := r.ReadCmd().Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestWriteCmd_ZeroBits_SetsWriteRegFlag(t *testing.T) {
	r := gorf24.Register(b("00000000"))
	expected := b("00100000")

	result := r.WriteCmd().Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestWriteCmd_AllOnes_MasksTwoMSBits(t *testing.T) {
	r := gorf24.Register(b("11111111"))
	expected := b("00111111")

	result := r.WriteCmd().Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}




















