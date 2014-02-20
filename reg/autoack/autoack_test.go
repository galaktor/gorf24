/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package autoack

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNew_RegisterAddress_IsEN_AA(t *testing.T) {
	a := New(0)
	expected := addr.EN_AA

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestNew_ReservedBitsSet_StoresZeroes(t *testing.T) {
	a := New(util.B("11111111"))
	expected := util.B("00111111")

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe0_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00111110"))
	expected := false

	actual := a.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe0_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00000001"))
	expected := true

	actual := a.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe1_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00111101"))
	expected := false

	actual := a.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe1_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00000010"))
	expected := true

	actual := a.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe2_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00111011"))
	expected := false

	actual := a.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe2_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00000100"))
	expected := true

	actual := a.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe3_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00110111"))
	expected := false

	actual := a.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe3_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00001000"))
	expected := true

	actual := a.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe4_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00101111"))
	expected := false

	actual := a.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe4_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00010000"))
	expected := true

	actual := a.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe5_Disabled_ReturnsFalse(t *testing.T) {
	a := New(util.B("00011111"))
	expected := false

	actual := a.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe5_Enabled_ReturnsTrue(t *testing.T) {
	a := New(util.B("00100000"))
	expected := true

	actual := a.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe6_InvalidButBitIsFlipped_ReturnsFalseAnyway(t *testing.T) {
	a := New(util.B("01000000"))
	expected := false

	actual := a.IsEnabled(pipe.P(6))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Enable_Pipe0_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00000001")

	a.Set(pipe.P0, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Enable_Pipe1_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00000010")

	a.Set(pipe.P1, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Enable_Pipe2_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00000100")

	a.Set(pipe.P2, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Enable_Pipe3_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00001000")

	a.Set(pipe.P3, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Enable_Pipe4_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00010000")

	a.Set(pipe.P4, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Enable_Pipe5_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00000000"))
	expected := util.B("00100000")

	a.Set(pipe.P5, true)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	} 
}

func TestSet_Disable_Pipe0_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00111110")

	a.Set(pipe.P0, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Disable_Pipe1_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00111101")

	a.Set(pipe.P1, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Disable_Pipe2_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00111011")

	a.Set(pipe.P2, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Disable_Pipe3_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00110111")

	a.Set(pipe.P3, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Disable_Pipe4_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00101111")

	a.Set(pipe.P4, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}

func TestSet_Disable_Pipe5_FlipsCorrectBit(t *testing.T) {
	a := New(util.B("00111111"))
	expected := util.B("00011111")

	a.Set(pipe.P5, false)

	actual := a.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with autoack '%v'", expected, actual, a)
	}
}
