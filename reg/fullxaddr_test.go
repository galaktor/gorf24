package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewFullXAddress_Pipe0_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewFullXAddress(pipe.P0, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullXAddress_Pipe1_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewFullXAddress(pipe.P1, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullXAddress_Pipe2_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewFullXAddress(pipe.P2, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullXAddress_Pipe3_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewFullXAddress(pipe.P3, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullXAddress_Pipe4_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewFullXAddress(pipe.P4, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullXAddress_Pipe5_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewFullXAddress(pipe.P5, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := XAddress(0)
	a := NewFullXAddress(pipe.P0, 0)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_NonZero_ReturnsRightBits(t *testing.T) {
	expected := XAddress(0xAFAFAFAFAF)
	a := NewFullXAddress(pipe.P0, 0xAFAFAFAFAF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_StoresRightBits(t *testing.T) {
	expected := XAddress(0xAFAFAFAFAF)
	a := NewFullXAddress(pipe.P0, XAddress(0))

	a.Set(0xAFAFAFAFAF)

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}
