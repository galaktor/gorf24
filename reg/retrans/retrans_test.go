/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package retrans

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_IsSETUP_RETR(t *testing.T) {
	r := New()
	expected := addr.SETUP_RETR

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestDisableCount_Enabled_SetsBitsToZero(t *testing.T) {
	r := NewWith(util.B("11111111"))
	expected := util.B("11110000")

	r.DisableCount()

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_Zero_FlipsRightBits(t *testing.T) {
	r := NewWith(util.B("11111111"))
	expected := util.B("11110000")

	err := r.SetCount(0)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_One_FlipsRightBits(t *testing.T) {
	r := NewWith(util.B("11110000"))
	expected := util.B("11110001")

	err := r.SetCount(1)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_One_ReturnsNoError(t *testing.T) {
	r := NewWith(util.B("11110000"))

	err := r.SetCount(1)

	if err != nil {
		t.Errorf("unexpected error '%v' with retrans '%v'", err, r)
	}
}

func TestSetCount_Two_FlipsRightBits(t *testing.T) {
	r := NewWith(util.B("11110000"))
	expected := util.B("11110010")

	err := r.SetCount(2)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_Fifteen_FlipsRightBits(t *testing.T) {
	r := NewWith(util.B("11110000"))
	expected := util.B("11111111")

	r.SetCount(15)

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_Sixteen_ReturnsError(t *testing.T) {
	r := NewWith(util.B("00000000"))
	expected := "value out of legal range: 16. Only values 0 - 15 allowed."
	
	err := r.SetCount(16)

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with retrans '%v'", expected, actual, r)
	}
}

func TestGetCount_Zero_ReturnsZero(t *testing.T) {
	r := NewWith(util.B("11110000"))
	expected := uint8(0)

	actual := r.GetCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestGetCount_One_ReturnsOne(t *testing.T) {
	r := NewWith(util.B("11110001"))
	expected := uint8(1)

	actual := r.GetCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestGetCount_Two_ReturnsTwo(t *testing.T) {
	r := NewWith(util.B("11110010"))
	expected := uint8(2)

	actual := r.GetCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestGetCount_Fifteen_ReturnsFifteen(t *testing.T) {
	r := NewWith(util.B("11111111"))
	expected := uint8(15)

	actual := r.GetCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestGetDelay_ZeroToFifteen_ReturnsCorrectDelay(t *testing.T) {
	for i := uint8(0); i <= 15; i++ {
		r := NewWith(i << 4) // iiii0000
		expected := Delay(i)

		actual := r.GetDelay()

		if actual != expected {
			t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
		}
	}
}

func TestSetDelay_ZeroToFifteen_FlipsRightBits(t *testing.T) {
	for i := uint8(0); i <= 15; i++ {
		r := NewWith(util.B("00000000"))
		expected := byte(i << 4)

		r.SetDelay(Delay(i))

		actual := r.Get()
		if actual != expected {
			t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
		}
	}
}
