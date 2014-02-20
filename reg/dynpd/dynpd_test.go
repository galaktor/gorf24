/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package dynpd

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
)

func TestNew_RegisterAddress_IsDYNPD(t *testing.T) {
	d := New(0)
	expected := addr.DYNPD

	actual := d.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestNew_ReservedBitsSet_StoresZeroes(t *testing.T) {
	d := New(util.B("11111111"))
	expected := util.B("00111111")

	actual := d.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe0_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00000001")

	err := d.Set(pipe.P0, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe1_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00000010")

	err := d.Set(pipe.P1, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe2_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00000100")

	err := d.Set(pipe.P2, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe3_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00001000")

	err := d.Set(pipe.P3, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe4_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00010000")

	err := d.Set(pipe.P4, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe5_Enable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00000000"))
	expected := util.B("00100000")

	err := d.Set(pipe.P5, true)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_PipeGreater5_Enable_ReturnsError(t *testing.T) {
	d := New(0)
	p := pipe.P(6)
	expected := "invalid pipe: 6"

	err := d.Set(p, true)

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe0_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00111110")

	err := d.Set(pipe.P0, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe1_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00111101")

	err := d.Set(pipe.P1, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe2_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00111011")

	err := d.Set(pipe.P2, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe3_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00110111")

	err := d.Set(pipe.P3, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe4_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00101111")

	err := d.Set(pipe.P4, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_Pipe5_Disable_FlipsRightBit(t *testing.T) {
	d := New(util.B("00111111"))
	expected := util.B("00011111")

	err := d.Set(pipe.P5, false)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	actual := d.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with dynpd '%v'", expected, actual, d)
	}
}

func TestSet_PipeGreater5_Disable_ReturnsError(t *testing.T) {
	d := New(0)
	p := pipe.P(6)
	expected := "invalid pipe: 6"

	err := d.Set(p, false)

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe0_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11111110"))
	expected := false

	actual := d.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe0_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00000001"))
	expected := true

	actual := d.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe1_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11111101"))
	expected := false

	actual := d.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe1_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00000010"))
	expected := true

	actual := d.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe2_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11111011"))
	expected := false

	actual := d.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe2_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00000100"))
	expected := true

	actual := d.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe3_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11110111"))
	expected := false

	actual := d.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe3_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00001000"))
	expected := true

	actual := d.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe4_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11101111"))
	expected := false

	actual := d.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe4_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00010000"))
	expected := true

	actual := d.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe5_BitZero_ReturnsFalse(t *testing.T) {
	d := New(util.B("11011111"))
	expected := false

	actual := d.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe5_BitOne_ReturnsFalse(t *testing.T) {
	d := New(util.B("00100000"))
	expected := true

	actual := d.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}

func TestIsEnabled_Pipe6_Invalid_ReturnsFalse(t *testing.T) {
	d := New(util.B("11111111"))
	expected := false

	actual := d.IsEnabled(pipe.P(6))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with dynpd '%v'", expected, actual, d)
	}
}
