package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
)


func TestIsEnabled_Pipe0_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11111110"))
	expected := false

	actual := a.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe0_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00000001"))
	expected := true

	actual := a.IsEnabled(pipe.P0)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe1_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11111101"))
	expected := false

	actual := a.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe1_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00000010"))
	expected := true

	actual := a.IsEnabled(pipe.P1)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe2_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11111011"))
	expected := false

	actual := a.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe2_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00000100"))
	expected := true

	actual := a.IsEnabled(pipe.P2)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe3_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11110111"))
	expected := false

	actual := a.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe3_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00001000"))
	expected := true

	actual := a.IsEnabled(pipe.P3)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe4_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11101111"))
	expected := false

	actual := a.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe4_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00010000"))
	expected := true

	actual := a.IsEnabled(pipe.P4)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe5_Disabled_ReturnsFalse(t *testing.T) {
	a := NewAutoAck(util.B("11011111"))
	expected := false

	actual := a.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe5_Enabled_ReturnsTrue(t *testing.T) {
	a := NewAutoAck(util.B("00100000"))
	expected := true

	actual := a.IsEnabled(pipe.P5)

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}

func TestIsEnabled_Pipe6_InvalidButBitIsFlipped_ReturnsFalseAnyway(t *testing.T) {
	a := NewAutoAck(util.B("01000000"))
	expected := false

	actual := a.IsEnabled(pipe.P(6))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with autoack '%v'", expected, actual, a)
	}
}
