package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

func someFullRxAddr(flags uint64) *FullXAddress {
	return NewFullRxAddress(pipe.P0, XAddress(flags))
}

func TestNewFullRxAddress_Pipe0_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewFullRxAddress(pipe.P0, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe1_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewFullRxAddress(pipe.P1, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe2_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewFullRxAddress(pipe.P2, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe3_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewFullRxAddress(pipe.P3, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe4_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewFullRxAddress(pipe.P4, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullRxAddress_Pipe5_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewFullRxAddress(pipe.P5, XAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}


func TestNewPartialRxAddress_Pipe0_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewPartialRxAddress(pipe.P0, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe1_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewPartialRxAddress(pipe.P1, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe2_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewPartialRxAddress(pipe.P2, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe3_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewPartialRxAddress(pipe.P3, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe4_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewPartialRxAddress(pipe.P4, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe5_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewPartialRxAddress(pipe.P5, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}