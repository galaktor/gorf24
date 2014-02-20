/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rfchan

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_IsRF_CH(t *testing.T) {
	c := New(0)
	expected := addr.RF_CH

	actual := c.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestNew_ReservedBitsOne_StoresZeroes(t *testing.T) {
	c := New(util.B("11111111"))
	expected := util.B("01111111")

	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestGet_Zeroes_ReturnsZero(t *testing.T) {
	c := New(util.B("00000000"))
	expected := uint8(0)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestGet_Ones_ReturnsMaxChannel(t *testing.T) {
	c := New(util.B("11111111"))
	expected := uint8(127)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestGet_FortyTwo_ReturnsFortyTwo(t *testing.T) {
	c := New(byte(42))
	expected := uint8(42)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_Zero_FlipsRightBits(t *testing.T) {
	c := New(util.B("11111111"))
	expected := util.B("00000000")

	err := c.Set(0)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_FortyTwo_FlipsRightBits(t *testing.T) {
	c := New(util.B("00000000"))
	expected := byte(42)

	err := c.Set(42)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_127_FlipsRightBits(t *testing.T) {
	c := New(util.B("00000000"))
	expected := byte(util.B("01111111"))

	err := c.Set(127)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_128_ReturnsError(t *testing.T) {
	c := New(util.B("00000000"))
	expected := "value outside of legal range: 128. Only values 1 - 127 allowed."

	err := c.Set(128)

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' on rfchan '%v'", expected, actual, c)
	}
}
