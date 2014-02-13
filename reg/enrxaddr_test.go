package reg

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/pipe"
)

func TestIsEnabled_Pipe0_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11111110"))
	expected := false

	actual := e.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe0_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00000001"))
	expected := true

	actual := e.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe1_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11111101"))
	expected := false

	actual := e.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe1_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00000010"))
	expected := true

	actual := e.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe2_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11111011"))
	expected := false

	actual := e.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe2_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00000100"))
	expected := true

	actual := e.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe3_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11110111"))
	expected := false

	actual := e.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe3_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00001000"))
	expected := true

	actual := e.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe4_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11101111"))
	expected := false

	actual := e.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe4_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00010000"))
	expected := true

	actual := e.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe5_FlagZero_ReturnsFalse(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("11011111"))
	expected := false

	actual := e.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}

func TestIsEnabled_Pipe5_FlagOne_ReturnsTrue(t *testing.T) {
	e := NewEnabledRxAddresses(util.B("00100000"))
	expected := true

	actual := e.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with enRxaddr '%v'", expected, actual, e)
	}
}









