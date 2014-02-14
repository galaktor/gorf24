package reg

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




















