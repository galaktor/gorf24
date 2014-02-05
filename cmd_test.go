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

func TestCMD_R_REGISTER_ZeroBits_ReturnsZero(t *testing.T) {
	r := gorf24.Register(gorf24.B("00000000"))
	expected := gorf24.B("00000000")

	result := gorf24.CMD_R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_R_REGISTER_AllOnes_ReturnsLastFiveBits(t *testing.T) {
	r := gorf24.Register(gorf24.B("11111111"))
	expected := gorf24.B("00011111")

	result := gorf24.CMD_R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_W_REGISTER_ZeroBits_SetsWriteRegFlag(t *testing.T) {
	r := gorf24.Register(gorf24.B("00000000"))
	expected := gorf24.B("00100000")

	result := gorf24.CMD_W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_W_REGISTER_AllOnes_MasksTwoMSBits(t *testing.T) {
	r := gorf24.Register(gorf24.B("11111111"))
	expected := gorf24.B("00111111")

	result := gorf24.CMD_W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}