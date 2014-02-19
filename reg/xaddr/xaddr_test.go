package xaddr

import (
	"testing"
)

func TestNewA_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := A(0x0000000000)
	
	actual := NewA(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewA_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := A(0xFFFFFFFFFF)
	
	actual := NewA(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewA_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := A(0xAAAAAAAAAA)
	
	actual := NewA(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewA_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := A(0xAAAAAAAAAA)

	actual := NewA(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := NewA(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















