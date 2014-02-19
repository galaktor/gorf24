/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
)

func TestGet_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := A(0)
	a := NewFull(addr.A(0), 0)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_NonZero_ReturnsRightBits(t *testing.T) {
	expected := A(0xAFAFAFAFAF)
	a := NewFull(addr.A(0), 0xAFAFAFAFAF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_StoresRightBits(t *testing.T) {
	expected := A(0xAFAFAFAFAF)
	a := NewFull(addr.A(0), A(0))

	a.Set(0xAFAFAFAFAF)

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}
