package util

import (
	"testing"
)

func TestB_Zeroes_ReturnsZero(t *testing.T) {
	expected := byte(0)
	
	actual := B("00000000")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB_Ones_ReturnsMaxByte(t *testing.T) {
	expected := byte(0xFF)
	
	actual := B("11111111")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB_MixedZeroesAndOnes_ReturnsRightNumber(t *testing.T) {
	expected := byte(0xAA)
	
	actual := B("10101010")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB4_Zeroes_ReturnsZero(t *testing.T) {
	expected := uint32(0)

	actual := B4("00000000000000000000000000000000")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB4_Ones_ReturnsMaxUint32(t *testing.T) {
	expected := uint32(0xFFFFFFFF)

	actual := B4("11111111111111111111111111111111")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB4_MixesZeroesAndOnes_ReturnsRightNumber(t *testing.T) {
	expected := uint32(0xAAAAAAAA)

	actual := B4("10101010101010101010101010101010")

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB5_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := [5]byte{0, 0, 0, 0, 0}

	actual := B5("0000000000000000000000000000000000000000")

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestB5_Ones_ReturnsMaxBytes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := B5("1111111111111111111111111111111111111111")

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestB5_MixedZeroesAndOnes_ReturnsRightBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	actual := B5("1010101010101010101010101010101010101010")

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestB5_LeadingZeroes_ReturnsRightBytes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0xAA, 0xAA, 0xAA}

	actual := B5("0000000000000000101010101010101010101010")

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestB5_TrailingZeroes_ReturnsRightBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0x00, 0x00}

	actual := B5("1010101010101010101010100000000000000000")

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestB5_Overflow5Bytes_ReturnsTruncated(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := B5("111111111111111111111111111111111111111111111111") // 6 bytes

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}





func TestItoB5_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := [5]byte{0, 0, 0, 0, 0}

	actual := ItoB5(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestItoB5_Ones_ReturnsMaxBytes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := ItoB5(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestItoB5_MixedZeroesAndOnes_ReturnsRightBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	actual := ItoB5(0xAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestItoB5_Overflow5Bytes_ReturnsTruncated(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := ItoB5(0xFFFFFFFFFFFF) // 6 bytes

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}






