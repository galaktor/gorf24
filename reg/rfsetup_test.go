package reg

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func NewRfSetup_RegisterAddress_IsRF_SETUP(t *testing.T) {
	s := NewRfSetup(0)
	expected := addr.RF_SETUP

	actual := s.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_Zero_ReturnsMin(t *testing.T) {
	s := NewRfSetup(util.B("11111001"))
	expected := PA_MIN
	
	actual := s.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_One_ReturnsLow(t *testing.T) {
	s := NewRfSetup(util.B("11111011"))
	expected := PA_LOW
	
	actual := s.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_Two_ReturnsMedium(t *testing.T) {
	s := NewRfSetup(util.B("11111101"))
	expected := PA_MEDIUM
	
	actual := s.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_Three_ReturnsMax(t *testing.T) {
	s := NewRfSetup(util.B("00000110"))
	expected := PA_MAX
	
	actual := s.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}
