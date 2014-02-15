package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

func someFullAddr(flags uint64) *FullPipeAddress {
	return NewFullPipeAddress(pipe.P0, PipeAddress(flags))
}

func TestNewPartialPipeAddress_Pipe0_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewPartialPipeAddress(pipe.P0, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialPipeAddress_Pipe1_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewPartialPipeAddress(pipe.P1, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialPipeAddress_Pipe2_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewPartialPipeAddress(pipe.P2, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialPipeAddress_Pipe3_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewPartialPipeAddress(pipe.P3, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialPipeAddress_Pipe4_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewPartialPipeAddress(pipe.P4, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartialPipeAddress_Pipe5_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewPartialPipeAddress(pipe.P5, someFullAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSBytesZero_FirstFourBytesZero(t *testing.T) {
	expected := PipeAddress(0x00000000FF)
	root := someFullAddr(0x0000000000)
	a := NewPartialPipeAddress(pipe.P0, root, 0xFF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSBytesOnes_FirstFourBytesOnes(t *testing.T) {
	expected := PipeAddress(0xFFFFFFFFAA)
	root := someFullAddr(0xFFFFFFFFFF)
	a := NewPartialPipeAddress(pipe.P0, root, 0xAA)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSByteChanges_ByteMsbChangesWithParent(t *testing.T) {
	expected := PipeAddress(0xFFFFFFFFA1)
	root := someFullAddr(0x0000000000)
	a := NewPartialPipeAddress(pipe.P0, root, 0xA1)
	
	root.Set(NewPipeAddress(0xFFFFFFFFFF))
	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}
