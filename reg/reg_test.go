/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"testing"
	
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_AddressFortyTwo_StoresThatAddress(t *testing.T) {
	expected := addr.A(42)
	r := New(expected, 0)

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func Test_FlagsFortyTwo_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), expected)

	actual := r.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestSet_FortyTwo_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), 0)

	r.Set(expected)

	actual := r.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}



















