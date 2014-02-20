/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package addrw

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_IsSETUP_AW(t *testing.T) {
	w := New(0)
	expected := addr.SETUP_AW

	actual := w.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestNew_ReservedBitsSet_StoresZeroes(t *testing.T) {
	w := New(util.B("11111111"))
	expected := util.B("00000011")

	actual := w.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_3Bytes_FlipsRightBit(t *testing.T) {
	w := New(util.B("00000000"))
	expected := util.B("00000001")

	w.Set(BYTES_3)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_4Bytes_FlipsRightBit(t *testing.T) {
	w := New(util.B("00000000"))
	expected := util.B("00000010")

	w.Set(BYTES_4)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_5Bytes_FlipsRightBit(t *testing.T) {
	w := New(util.B("00000000"))
	expected := util.B("00000011")

	w.Set(BYTES_5)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_GreaterThree_ReturnsError(t *testing.T) {
	w := New(util.B("00000000"))
	expected := "value outside of legal range: 4. only values 1 - 3 allowed."
	
	err := w.Set(Width(4))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_Zero_ReturnsError(t *testing.T) {
	w := New(util.B("00000000"))
	expected := "value outside of legal range: 0. only values 1 - 3 allowed."
	
	err := w.Set(Width(0))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueOne_Returns3Bytes(t *testing.T) {
	w := New(util.B("00000001"))
	expected := BYTES_3

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueTwo_Returns4Bytes(t *testing.T) {
	w := New(util.B("00000010"))
	expected := BYTES_4

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueThree_Returns5Bytes(t *testing.T) {
	w := New(util.B("00000011"))
	expected := BYTES_5

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}
