/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNew_AddressFortyTwo_StoresThatAddress(t *testing.T) {
	expected := addr.A(42)
	r := New(expected, 0)

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestNew_MaskFortyTwo_StoresThaMask(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), expected)

	actual := r.Mask()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestNO_MASK_Is_AllOnes(t *testing.T) {
	expected := util.B("11111111")

	actual := NO_MASK

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestSet_FortyTwo_NoMask_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)

	r.Set(expected)

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestSet_FortyTwo_ZeroMask_StoresZero(t *testing.T) {
	expected := byte(0)
	r := New(addr.A(0), 0x00)

	r.Set(expected)

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestSet_AllOnes_MixedMask_StoresMaskedValue(t *testing.T) {
	expected := byte(0xAA)
	r := New(addr.A(0), 0xAA)

	r.Set(util.B("11111111"))

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestGet_FortyTwo_NoMask_ReturnsFortyTwo(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)
	r.flags = 42

	actual := r.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_OneByte_NoMask_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)

	r.Write([]byte{expected})

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_OneByte_Masked_StoresMaskedByte(t *testing.T) {
	expected := byte(0xAA)
	r := New(addr.A(0), 0xAA)

	r.Write([]byte{util.B("11111111")})

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_OneByte_ReturnsNilError(t *testing.T) {
	r := New(addr.A(0), NO_MASK)

	_,err := r.Write([]byte{0})

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}
}

func TestWrite_OneByte_ReturnsOne(t *testing.T) {
	expected := 1
	r := New(addr.A(0), NO_MASK)

	actual,_ := r.Write([]byte{0})

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoBytes_DoesNotChangeFlags(t *testing.T) {
	expected := util.B("11111111")
	r := New(addr.A(0), NO_MASK)
	r.Set(expected)

	r.Write(make([]byte, 0))

	actual := r.flags
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoByte_ReturnsError(t *testing.T) {
	expected := "cannot write 0 bytes. register can only store exactly 1 byte"
	r := New(addr.A(0), NO_MASK)

	_, err := r.Write(make([]byte, 0))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_TwoBytes_ReturnsError(t *testing.T) {
	expected := "cannot write 2 bytes. register can only store exactly 1 byte"
	r := New(addr.A(0), NO_MASK)

	_, err := r.Write(make([]byte, 2))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoByte_ReturnsZero(t *testing.T) {
	expected := 0
	r := New(addr.A(0), NO_MASK)

	actual, _ := r.Write(make([]byte, 0))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_TwoBytes_ReturnsZero(t *testing.T) {
	expected := 0
	r := New(addr.A(0), NO_MASK)

	actual, _ := r.Write(make([]byte, 2))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}
