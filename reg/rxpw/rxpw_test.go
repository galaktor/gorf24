/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rxpw

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_AllPipes_HasRightRegister(t *testing.T) {
	for i := 0; i <= 5; i++ {
		p := pipe.P(i)
		w := New(p, 0)
		expected := addr.RX_PW(p)
		
		actual := w.Address()

		if actual != expected {
			t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
		}
	}
}

func TestNew_ReservedBitsOne_StoresZeroes(t *testing.T) {
	w := New(pipe.P0, util.B("11111111"))
	expected := util.B("00111111")

	actual := w.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
	}
}

func TestSet_Zero_FlipsRelevantBits(t *testing.T) {
	w := New(pipe.P0, util.B("00111111"))
	expected := util.B("00000000")

	w.Set(0)
	
	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
	}
}

func TestSet_One_FlipsRelevantBits(t *testing.T) {
	w := New(pipe.P0, util.B("0000000"))
	expected := util.B("00000001")

	w.Set(1)
	
	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
	}
}

func TestSet_ThirtyTwo_FlipsRelevantBits(t *testing.T) {
	w := New(pipe.P0, util.B("00000000"))
	expected := util.B("00100000")

	w.Set(32)
	
	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
	}
}

func TestSet_OverThrityTwo_ReturnsError(t *testing.T) {
	w := New(pipe.P0, util.B("00000000"))
	expected := "value out of legal range: 33. allowed values from 0 - 32"

	err := w.Set(33)
	
	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rxpw '%v'", expected, actual, w)
	}
}

func TestGet_Zero_ReturnsZero(t *testing.T) {
	w := New(pipe.P0, util.B("00000000"))
	expected := uint8(0)

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rxpw '%v'", expected, actual, w)
	}
}

func TestGet_ThirtyTwo_ReturnsThirtyTwo(t *testing.T) {
	w := New(pipe.P0, util.B("00100000"))
	expected := uint8(32)

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rxpw '%v'", expected, actual, w)
	}
}

func TestGet_OverThirtyTwo_TruncatesAndReturnsThirtyTwo(t *testing.T) {
	w := New(pipe.P0, byte(33))
	expected := uint8(32)

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rxpw '%v'", expected, actual, w)
	}
}
