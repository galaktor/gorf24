package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewFullRxAddress_Pipe0_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewFullRxAddress(pipe.P0, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe1_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewFullRxAddress(pipe.P1, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe2_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewFullRxAddress(pipe.P2, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe3_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewFullRxAddress(pipe.P3, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe4_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewFullRxAddress(pipe.P4, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe5_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewFullRxAddress(pipe.P5, [5]byte{})

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestNewFullRxAddress_FlagsParam_IsCopiedNotReferenced(t *testing.T) {
	provided := NewRxAddress(0x00000000, 0x00)
	a := NewFullRxAddress(pipe.P0, provided)

	if &provided == &a.flags {
		t.Errorf("expected different address but found same as '%p' for fulladr '%v'", &provided, a)
	}
}

func TestByte_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := NewRxAddress(0x00000000, 0x00)
	a := NewFullRxAddress(pipe.P0, expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' for fulladdr '%v'", expected, actual, a)
	}
}

func TestByte_NonZero_ReturnsRightBits(t *testing.T) {
	expected := NewRxAddress(0xAFAFAF01, 0xA1)
	a := NewFullRxAddress(pipe.P0, expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_ByteReturnsRightBits(t *testing.T) {
	expected := NewRxAddress(0xAFAFAF01, 0xA1)
	a := NewFullRxAddress(pipe.P0, NewRxAddress(0x0, 0x0))

	a.Set(expected)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' for fulladdr '%v'", expected, actual, a)
	}
}
