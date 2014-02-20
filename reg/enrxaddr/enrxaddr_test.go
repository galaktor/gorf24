/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package enrxaddr

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

func TestNew_RegisterAddress_IsEN_RXADDR(t *testing.T) {
	e := New(0)
	expected := addr.EN_RXADDR

	actual := e.Address()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestNew_ReservedBitsOne_StoresZeroes(t *testing.T) {
	e := New(util.B("11111111"))
	expected := util.B("00111111")

	actual := e.Byte()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe0_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11111110"))
	expected := false

	actual := e.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe0_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00000001"))
	expected := true

	actual := e.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe1_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11111101"))
	expected := false

	actual := e.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe1_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00000010"))
	expected := true

	actual := e.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe2_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11111011"))
	expected := false

	actual := e.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe2_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00000100"))
	expected := true

	actual := e.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe3_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11110111"))
	expected := false

	actual := e.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe3_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00001000"))
	expected := true

	actual := e.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe4_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11101111"))
	expected := false

	actual := e.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe4_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00010000"))
	expected := true

	actual := e.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe5_FlagZero_ReturnsFalse(t *testing.T) {
	e := New(util.B("11011111"))
	expected := false

	actual := e.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe5_FlagOne_ReturnsTrue(t *testing.T) {
	e := New(util.B("00100000"))
	expected := true

	actual := e.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe0_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00000001")

	e.Set(pipe.P0, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe0_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00111110")

	e.Set(pipe.P0, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe1_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00000010")

	e.Set(pipe.P1, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe1_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00111101")

	e.Set(pipe.P1, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe2_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00000100")

	e.Set(pipe.P2, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe2_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00111011")

	e.Set(pipe.P2, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe3_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00001000")

	e.Set(pipe.P3, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe3_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00110111")

	e.Set(pipe.P3, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe4_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00010000")

	e.Set(pipe.P4, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe4_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00101111")

	e.Set(pipe.P4, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Enable_Pipe5_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00000000"))
	expected := util.B("00100000")

	e.Set(pipe.P5, true)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestSet_Disable_Pipe5_FlipsCorrectBit(t *testing.T) {
	e := New(util.B("00111111"))
	expected := util.B("00011111")

	e.Set(pipe.P5, false)

	actual := e.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with enRxaddr '%v'", expected, actual, e)
	}
}










