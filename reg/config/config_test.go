/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package config

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNew_RegisterAddress_IsCONFIG(t *testing.T) {
	c := New()
	expected := addr.CONFIG

	actual := c.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestNew_ReservedBitsSet_StoresZeroes(t *testing.T) {
	c := NewWith(util.B("11111111"))
	expected := util.B("01111111")
	
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryReceiver_RelevantFlagZero_SetsFlagToOne(t *testing.T) {
	c := NewWith(util.B("00000000"))
	expected := util.B("00000001")

	c.SetPrimaryReceiver()

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryReceiver_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("00101010"))
	expected := util.B("00101011")

	c.SetPrimaryReceiver()

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagZero_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("01111110"))
	expected := false

	actual := c.IsPrimaryReceiver()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagOne_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("00000001"))
	expected := true

	actual := c.IsPrimaryReceiver()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_SetsBitToZero(t *testing.T) {
	c := NewWith(util.B("00000001"))
	expected := util.B("00000000")

	c.SetPrimaryTransmitter()

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_DoesntFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("00101010"))
	expected := util.B("00101010")

	c.SetPrimaryTransmitter()

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagZero_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("01111110"))
	expected := true

	actual := c.IsPrimaryTransmitter()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagOne_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("00000001"))
	expected := false

	actual := c.IsPrimaryTransmitter()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_True_SetsFlagToOne(t *testing.T) {
	c := NewWith(util.B("01111101"))
	expected := util.B("01111111")

	c.SetPowerUp(true)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_False_SetsFlagToZero(t *testing.T) {
	c := NewWith(util.B("00000010"))
	expected := util.B("00000000")

	c.SetPowerUp(false)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("01010101"))

	expected := util.B("01010111")
	c.SetPowerUp(true)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetPowerUp(false)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagZero_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("01111101"))
	expected := false

	actual := c.IsPowerUp()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagOne_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("00000010"))
	expected := true

	actual := c.IsPowerUp()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To16bits_SetsFlagToOne(t *testing.T) {
	c := NewWith(util.B("01111011"))
	expected := util.B("01111111")

	c.SetCrcLength(CRC_16BIT)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("01010101"))

	expected := util.B("01010001")
	c.SetCrcLength(CRC_8BIT)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetCrcLength(CRC_16BIT)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To8bits_SetsFlagToZero(t *testing.T) {
	c := NewWith(util.B("00000100"))
	expected := util.B("00000000")

	c.SetCrcLength(CRC_8BIT)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_8bits_Returns8bits(t *testing.T) {
	c := New()
	expected := CRC_8BIT
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_16bits_Returns16bits(t *testing.T) {
	c := New()
	expected := CRC_16BIT
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_True_SetsFlagToOne(t *testing.T) {
	c := NewWith(util.B("01110111"))
	expected := util.B("01111111")

	c.SetCrcEnabled(true)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_False_SetsFlagToZero(t *testing.T) {
	c := NewWith(util.B("00001000"))
	expected := util.B("00000000")

	c.SetCrcEnabled(false)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("01010101"))

	expected := util.B("01011101")
	c.SetCrcEnabled(true)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetCrcEnabled(false)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b'but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsCrcEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("01110111"))
	expected := false

	actual := c.IsCrcEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsCrcEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("00001000"))
	expected := true

	actual := c.IsCrcEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := NewWith(util.B("00010000"))
	expected := util.B("00000000")

	c.SetMaxRtInterruptEnabled(true)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := NewWith(util.B("01101111"))
	expected := util.B("01111111")

	c.SetMaxRtInterruptEnabled(false)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("00101010"))

	expected := util.B("00111010")
	c.SetMaxRtInterruptEnabled(false)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("00101010")
	c.SetMaxRtInterruptEnabled(true)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsMaxRtInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("00010000"))
	expected := false

	actual := c.IsMaxRtInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsMaxRtInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("01101111"))
	expected := true

	actual := c.IsMaxRtInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := NewWith(util.B("00100000"))
	expected := util.B("00000000")

	c.SetTxDsInterruptEnabled(true)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := NewWith(util.B("01011111"))
	expected := util.B("01111111")

	c.SetTxDsInterruptEnabled(false)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("01010101"))

	expected := util.B("01110101")
	c.SetTxDsInterruptEnabled(false)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetTxDsInterruptEnabled(true)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsTxDsInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("00100000"))
	expected := false

	actual := c.IsTxDsInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsTxDsInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("01011111"))
	expected := true

	actual := c.IsTxDsInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := NewWith(util.B("01000000"))
	expected := util.B("00000000")

	c.SetRxDrInterruptEnabled(true)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := NewWith(util.B("00111111"))
	expected := util.B("01111111")

	c.SetRxDrInterruptEnabled(false)

	actual := c.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := NewWith(util.B("00101010"))

	expected := util.B("01101010")
	c.SetRxDrInterruptEnabled(false)
	actual := c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("00101010")
	c.SetRxDrInterruptEnabled(true)
	actual = c.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsRxDrInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := NewWith(util.B("01000000"))
	expected := false

	actual := c.IsRxDrInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsRxDrInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := NewWith(util.B("00111111"))
	expected := true

	actual := c.IsRxDrInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}