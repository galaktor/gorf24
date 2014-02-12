package reg

import (
	"testing"
	
	"github.com/galaktor/gorf24/util"
)

func TestNewRxAddress_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := RxAddress(util.ItoB5(0x0))
	
	actual := NewRxAddress(0x0, 0x0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := RxAddress(util.ItoB5(0xFFFFFFFFFF))
	
	actual := NewRxAddress(0xFFFFFFFF, 0xFF)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_MsbAllOnes_FirstFourBytesAllOnes(t *testing.T) {
	expected := RxAddress(util.ItoB5(0xFFFFFFFF00))
	
	actual := NewRxAddress(0xFFFFFFFF, 0x00)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_LsbAllOnes_LastByteAllOnes(t *testing.T) {
	expected := RxAddress(util.ItoB5(0x00000000FF))
	
	actual := NewRxAddress(0x00000000, 0xFF)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_MsbMixedOnesAndZeroes_FirstFourBytesMixed(t *testing.T) {
	expected := RxAddress(util.ItoB5(0xAAAAAAAA00))
	
	actual := NewRxAddress(0xAAAAAAAA, 0x00)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_LsbMixedOnesAndZeroes_LastByteMixed(t *testing.T) {
	expected := RxAddress(util.ItoB5(0x00000000AA))
	
	actual := NewRxAddress(0x00000000, 0xAA)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_MsbLeadingZeroes_ReturnsRightMsb(t *testing.T) {
	expected := RxAddress(util.ItoB5(0x0000AAAA00))
	
	actual := NewRxAddress(0x0000AAAA, 0x00)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewRxAddress_LsbLeadingZeroes_ReturnsRightLsb(t *testing.T) {
	expected := RxAddress(util.ItoB5(0xFFFFFFFF0A))
	
	actual := NewRxAddress(0xFFFFFFFF, 0x0A)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}
