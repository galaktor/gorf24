package xaddr

import (
	"testing"
)

func TestNew_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := A(0x0000000000)
	
	actual := New(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := A(0xFFFFFFFFFF)
	
	actual := New(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := A(0xAAAAAAAAAA)
	
	actual := New(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := A(0xAAAAAAAAAA)

	actual := New(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := New(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















