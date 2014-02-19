/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rpd

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNew_RegisterAddress_IsRPD(t *testing.T) {
	r := New(0)
	expected := addr.RPD

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rpd '%v'", expected, actual, r)
	}
}

func TestTriggered_BitZero_ReturnsFalse(t *testing.T) {
	r := New(util.B("11111110"))
	expected := false

	actual := r.Triggered()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rpd '%v'", expected, actual, r)
	}
}

func TestTriggered_BitOne_ReturnsTrue(t *testing.T) {
	r := New(util.B("00000001"))
	expected := true

	actual := r.Triggered()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rpd '%v'", expected, actual, r)
	}
}
