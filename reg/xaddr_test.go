package reg

import (
	"testing"
)

func TestNewXAddress_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := XAddress(0x0000000000)
	
	actual := NewXAddress(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewXAddress_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := XAddress(0xFFFFFFFFFF)
	
	actual := NewXAddress(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewXAddress_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := XAddress(0xAAAAAAAAAA)
	
	actual := NewXAddress(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewXAddress_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := XAddress(0xAAAAAAAAAA)

	actual := NewXAddress(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := NewXAddress(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















