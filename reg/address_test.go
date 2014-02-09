package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
)

func TestRxAddr_Pipe0_ReturnsRightAddress(t *testing.T) {
	p := pipe.P0
	expected := Address(0x0A)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxAddr_Pipe1_ReturnsRightAddress(t *testing.T) {
	p := pipe.P1
	expected := Address(0x0B)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxAddr_Pipe2_ReturnsRightAddress(t *testing.T) {
	p := pipe.P2
	expected := Address(0x0C)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxAddr_Pipe3_ReturnsRightAddress(t *testing.T) {
	p := pipe.P3
	expected := Address(0x0D)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxAddr_Pipe4_ReturnsRightAddress(t *testing.T) {
	p := pipe.P4
	expected := Address(0x0E)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxAddr_Pipe5_ReturnsRightAddress(t *testing.T) {
	p := pipe.P5
	expected := Address(0x0F)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

// TODO; should probably panic
func TestRxAddr_PipeOver5_ReturnsSixteenEvenIfThatsInvalid(t *testing.T) {
	p := pipe.P(6)
	expected := Address(0x10)

	actual := RX_ADDR(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe0_ReturnsRightAddress(t *testing.T) {
	p := pipe.P0
	expected := Address(0x11)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe1_ReturnsRightAddress(t *testing.T) {
	p := pipe.P1
	expected := Address(0x12)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe2_ReturnsRightAddress(t *testing.T) {
	p := pipe.P2
	expected := Address(0x13)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe3_ReturnsRightAddress(t *testing.T) {
	p := pipe.P3
	expected := Address(0x14)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe4_ReturnsRightAddress(t *testing.T) {
	p := pipe.P4
	expected := Address(0x15)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}

func TestRxPw_Pipe5_ReturnsRightAddress(t *testing.T) {
	p := pipe.P5
	expected := Address(0x16)

	actual := RX_PW(p)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, actual, p)
	}
}