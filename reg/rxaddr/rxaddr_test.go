/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rxaddr

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/reg/xaddr"
	"github.com/galaktor/gorf24/pipe"
)

func someFullRxAddr(flags uint64) *xaddr.Full {
	return NewFull(pipe.P0, xaddr.A(flags))
}

func TestNewFull_Pipe0_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewFull(pipe.P0, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFull_Pipe1_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewFull(pipe.P1, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFull_Pipe2_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewFull(pipe.P2, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFull_Pipe3_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewFull(pipe.P3, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFull_Pipe4_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewFull(pipe.P4, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
	
}

func TestNewFull_Pipe5_HasP0RegisterAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewFull(pipe.P5, xaddr.A(0))

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}


func TestNewPartial_Pipe0_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)
	a := NewPartial(pipe.P0, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartial_Pipe1_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)
	a := NewPartial(pipe.P1, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartial_Pipe2_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)
	a := NewPartial(pipe.P2, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartial_Pipe3_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)
	a := NewPartial(pipe.P3, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartial_Pipe4_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)
	a := NewPartial(pipe.P4, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestNewPartial_Pipe5_HasRightRegAddress(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)
	a := NewPartial(pipe.P5, someFullRxAddr(0), 0x0)

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}