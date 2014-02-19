package rpd

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNewRPD_RegisterAddress_IsRPD(t *testing.T) {
	r := NewRPD(0)
	expected := addr.RPD

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rpd '%v'", expected, actual, r)
	}
}

func TestTriggered_BitZero_ReturnsFalse(t *testing.T) {
	r := NewRPD(util.B("11111110"))
	expected := false

	actual := r.Triggered()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rpd '%v'", expected, actual, r)
	}
}

func TestTriggered_BitOne_ReturnsTrue(t *testing.T) {
	r := NewRPD(util.B("00000001"))
	expected := true

	actual := r.Triggered()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rpd '%v'", expected, actual, r)
	}
}
