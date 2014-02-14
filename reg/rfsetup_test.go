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

func TestSetPower_Min_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111001")
	
	s.Set(PA_MIN)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Low_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111011")
	
	s.Set(PA_LOW)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Medium_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111101")
	
	s.Set(PA_MEDIUM)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Max_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111111")
	
	s.Set(PA_MAX)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_GreaterThanThree_ReturnsError(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := "Value out of legal range: 4. Allowed value from 0 -3."
	
	err := s.Set(PowerLevel(4))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rfsetup '%v'", expected, actual, s)
	}
}




















