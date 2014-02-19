package txobs

import (
	"testing"
	
	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_IsOBSERVE_TX(t *testing.T) {
	o := New(0)
	expected := addr.OBSERVE_TX

	actual := o.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestRetransmittedPacketCount_Zero_ReturnsZero(t *testing.T) {
	o := New(util.B("11110000"))
	expected := uint8(0)

	actual := o.RetransmittedPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestRetransmittedPacketCount_AllOnes_Returns31(t *testing.T) {
	o := New(util.B("00001111"))
	expected := uint8(15)

	actual := o.RetransmittedPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestRetransmittedPacketCount_Seven_Returns7(t *testing.T) {
	o := New(byte(7))
	expected := uint8(7)

	actual := o.RetransmittedPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestLostPacketCount_Zero_ReturnsZero(t *testing.T) {
	o := New(util.B("00001111"))
	expected := uint8(0)

	actual := o.LostPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestLostPacketCount_AllOnes_ReturnsZero(t *testing.T) {
	o := New(util.B("11110000"))
	expected := uint8(15)

	actual := o.LostPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}

func TestLostPacketCount_Seven_ReturnsSeven(t *testing.T) {
	o := New(util.B("01110000"))
	expected := uint8(7)

	actual := o.LostPacketCount()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}
