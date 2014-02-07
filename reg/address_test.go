package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
)

func TestRegRxAddr_Pipe0_ReturnsRightRegister(t *testing.T) {
	p := pipe.P0
	expected := Address(0x0A)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxAddr_Pipe1_ReturnsRightRegister(t *testing.T) {
	p := pipe.P1
	expected := Address(0x0B)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxAddr_Pipe2_ReturnsRightRegister(t *testing.T) {
	p := pipe.P2
	expected := Address(0x0C)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxAddr_Pipe3_ReturnsRightRegister(t *testing.T) {
	p := pipe.P3
	expected := Address(0x0D)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxAddr_Pipe4_ReturnsRightRegister(t *testing.T) {
	p := pipe.P4
	expected := Address(0x0E)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxAddr_Pipe5_ReturnsRightRegister(t *testing.T) {
	p := pipe.P5
	expected := Address(0x0F)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

// TODO; should probably panic
func TestRegRxAddr_PipeOver5_ReturnsSixteenEvenIfThatsInvalid(t *testing.T) {
	p := pipe.P(6)
	expected := Address(0x10)

	actual := REG_RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe0_ReturnsRightRegister(t *testing.T) {
	p := pipe.P0
	expected := Address(0x11)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe1_ReturnsRightRegister(t *testing.T) {
	p := pipe.P1
	expected := Address(0x12)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe2_ReturnsRightRegister(t *testing.T) {
	p := pipe.P2
	expected := Address(0x13)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe3_ReturnsRightRegister(t *testing.T) {
	p := pipe.P3
	expected := Address(0x14)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe4_ReturnsRightRegister(t *testing.T) {
	p := pipe.P4
	expected := Address(0x15)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRegRxPw_Pipe5_ReturnsRightRegister(t *testing.T) {
	p := pipe.P5
	expected := Address(0x16)

	actual := REG_RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}