package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewFullPipeAddress_Pipe0_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewFullPipeAddress(pipe.P0, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullPipeAddress_Pipe1_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewFullPipeAddress(pipe.P1, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullPipeAddress_Pipe2_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewFullPipeAddress(pipe.P2, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullPipeAddress_Pipe3_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewFullPipeAddress(pipe.P3, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullPipeAddress_Pipe4_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewFullPipeAddress(pipe.P4, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFullPipeAddress_Pipe5_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewFullPipeAddress(pipe.P5, PipeAddress(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := PipeAddress(0)
	a := NewFullPipeAddress(pipe.P0, 0)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_NonZero_ReturnsRightBits(t *testing.T) {
	expected := PipeAddress(0xAFAFAFAFAF)
	a := NewFullPipeAddress(pipe.P0, 0xAFAFAFAFAF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_StoresRightBits(t *testing.T) {
	expected := PipeAddress(0xAFAFAFAFAF)
	a := NewFullPipeAddress(pipe.P0, PipeAddress(0))

	a.Set(0xAFAFAFAFAF)

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}
