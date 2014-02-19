package rfchan

import (
	"testing"

	"github.com/galaktor/gorf24/util"
)

func TestGet_Zeroes_ReturnsZero(t *testing.T) {
	c := NewRfChannel(util.B("00000000"))
	expected := uint8(0)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestGet_Ones_ReturnsMaxChannel(t *testing.T) {
	c := NewRfChannel(util.B("11111111"))
	expected := uint8(127)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestGet_FortyTwo_ReturnsFortyTwo(t *testing.T) {
	c := NewRfChannel(byte(42))
	expected := uint8(42)

	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_Zero_FlipsRightBits(t *testing.T) {
	c := NewRfChannel(util.B("11111111"))
	expected := util.B("00000000")

	err := c.Set(0)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_FortyTwo_FlipsRightBits(t *testing.T) {
	c := NewRfChannel(util.B("00000000"))
	expected := byte(42)

	err := c.Set(42)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_127_FlipsRightBits(t *testing.T) {
	c := NewRfChannel(util.B("00000000"))
	expected := byte(util.B("01111111"))

	err := c.Set(127)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rfchan '%v'", expected, actual, c)
	}
}

func TestSet_128_ReturnsError(t *testing.T) {
	c := NewRfChannel(util.B("00000000"))
	expected := "value outside of legal range: 128. Only values 1 - 127 allowed."

	err := c.Set(128)

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' on rfchan '%v'", expected, actual, c)
	}
}
