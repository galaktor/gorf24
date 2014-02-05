package gorf24

import (
	"testing"

	"."
)

func TestB_ZeroByte_ReturnsZero(t *testing.T) {
	bits := "00000000"
	expected := byte(0)
	
	result := gorf24.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with bits '%s'", expected, result, bits)
	}
}

func TestB_AllOnes_Returns255(t *testing.T) {
	bits := "11111111"
	expected := byte(255)
	
	result := gorf24.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}

func TestB_OverflowWithNineBits_AllOnes_Returns255(t *testing.T) {
	bits := "111111111"
	expected := byte(255)
	
	result := gorf24.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}

func TestB_FortyTwo_ReturnsFortyTwo(t *testing.T) {
	bits := "101010"
	expected := byte(42)
	
	result := gorf24.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}
