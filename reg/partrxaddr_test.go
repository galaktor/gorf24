package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

func someFullAddr(msb uint32, lsb byte) *FullRxAddress {
	return NewFullRxAddress(pipe.P0, NewRxAddress(msb, lsb))
}

func TestNewPartialRxAddress_Pipe0_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewPartialRxAddress(pipe.P0, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe1_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewPartialRxAddress(pipe.P1, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe2_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewPartialRxAddress(pipe.P2, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe3_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewPartialRxAddress(pipe.P3, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe4_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewPartialRxAddress(pipe.P4, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialRxAddress_Pipe5_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewPartialRxAddress(pipe.P5, someFullAddr(0,0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSBytesZero_FirstFourBytesZero(t *testing.T) {
	expected := NewRxAddress(0x00000000, 0x00)
	root := someFullAddr(0x00000000, 0xFF)
	a := NewPartialRxAddress(pipe.P0, root, 0x00)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSBytesOnes_FirstFourBytesOnes(t *testing.T) {
	expected := NewRxAddress(0xFFFFFFFF, 0xFF)
	root := someFullAddr(0xFFFFFFFF, 0x00)
	a := NewPartialRxAddress(pipe.P0, root, 0xFF)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_LSByteNonZero_ParentLSByteZero_LastByteMatches(t *testing.T) {
	expected := NewRxAddress(0x00000000, 0xA1)
	root := someFullAddr(0x00000000, 0x00)
	a := NewPartialRxAddress(pipe.P0, root, 0xA1)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSByteChanges_ByteMsbChangesWithParent(t *testing.T) {
	expected := NewRxAddress(0xFFFFFFFF, 0xA1)
	root := someFullAddr(0x00000000, 0x00)
	a := NewPartialRxAddress(pipe.P0, root, 0xA1)
	
	root.Set(NewRxAddress(0xFFFFFFFF, 0xFF))
	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}













