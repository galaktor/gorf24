package reg

import (
	"testing"
)

func TestNewRxAddress_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := RxAddress(0x0000000000)
	
	actual := NewRxAddress(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewRxAddress_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := RxAddress(0xFFFFFFFFFF)
	
	actual := NewRxAddress(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewRxAddress_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := RxAddress(0xAAAAAAAAAA)
	
	actual := NewRxAddress(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewRxAddress_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := RxAddress(0xAAAAAAAAAA)

	actual := NewRxAddress(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := NewRxAddress(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















