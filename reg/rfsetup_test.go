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
	
	actual := s.GetPowerLevel()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_One_ReturnsLow(t *testing.T) {
	s := NewRfSetup(util.B("11111011"))
	expected := PA_LOW
	
	actual := s.GetPowerLevel()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_Two_ReturnsMedium(t *testing.T) {
	s := NewRfSetup(util.B("11111101"))
	expected := PA_MEDIUM
	
	actual := s.GetPowerLevel()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetPower_Three_ReturnsMax(t *testing.T) {
	s := NewRfSetup(util.B("00000110"))
	expected := PA_MAX
	
	actual := s.GetPowerLevel()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Min_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111001")
	
	s.SetPowerLevel(PA_MIN)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Low_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111011")
	
	s.SetPowerLevel(PA_LOW)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Medium_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111101")
	
	s.SetPowerLevel(PA_MEDIUM)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_Max_FlipsRightBits(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11111111")
	
	s.SetPowerLevel(PA_MAX)

	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPower_GreaterThanThree_ReturnsError(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := "Value out of legal range: 4. Allowed value from 0 -3."
	
	err := s.SetPowerLevel(PowerLevel(4))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetDatarate_DrLowBitZero_DrHighBitZero_Returns1Mbps(t *testing.T) {
	s := NewRfSetup(util.B("11010111"))
	expected := RATE_1MBPS
	
	actual := s.GetDatarate()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetDatarate_DrLowBitZero_DrHighBitOne_Returns2Mbps(t *testing.T) {
	s := NewRfSetup(util.B("11011111"))
	expected := RATE_2MBPS
	
	actual := s.GetDatarate()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetDatarate_DrLowBitOne_DrHighBitZero_Returns250Kbps(t *testing.T) {
	s := NewRfSetup(util.B("11110111"))
	expected := RATE_250KBPS
	
	actual := s.GetDatarate()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestGetDatarate_DrLowBitOne_DrHighBitOne_Returns250Kbps(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := RATE_250KBPS
	
	actual := s.GetDatarate()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestIsPllLockEnabled_Zero_ReturnsFalse(t *testing.T) {
	s := NewRfSetup(util.B("11101111"))
	expected := false
	
	actual := s.IsPllLockEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rfsetup '%v'", expected, actual, s)
	}
}

func TestIsPllLockEnabled_One_ReturnsTrue(t *testing.T) {
	s := NewRfSetup(util.B("00010000"))
	expected := true
	
	actual := s.IsPllLockEnabled()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPllLock_Enabled_BitIsOne(t *testing.T) {
	s := NewRfSetup(util.B("00000000"))
	expected := util.B("00010000")
	
	s.SetPllLock(true)
	
	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}

func TestSetPllLock_Disabled_BitIsOne(t *testing.T) {
	s := NewRfSetup(util.B("11111111"))
	expected := util.B("11101111")
	
	s.SetPllLock(false)
	
	actual := s.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with rfsetup '%v'", expected, actual, s)
	}
}










