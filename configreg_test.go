package gorf24

import (
	"testing"

	"."
)

func TestSetPrimaryReceiver_RelevantFlagZero_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000000"))
	expected := gorf24.B("00000001")

	c.SetPrimaryReceiver()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryReceiver_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101011")

	c.SetPrimaryReceiver()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagZero_ReturnsFalse(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111110"))
	expected := false

	actual := c.IsPrimaryReceiver()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := true

	actual := c.IsPrimaryReceiver()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_SetsBitToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := gorf24.B("00000000")

	c.SetPrimaryTransmitter()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_DoesntFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101010")

	c.SetPrimaryTransmitter()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagZero_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111110"))
	expected := true

	actual := c.IsPrimaryTransmitter()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := false

	actual := c.IsPrimaryTransmitter()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_True_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111101"))
	expected := gorf24.B("11111111")

	c.SetPowerUp(true)
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_False_SetsFlagToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000010"))
	expected := gorf24.B("00000000")

	c.SetPowerUp(false)
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("01010101"))

	expected := gorf24.B("01010111")
	c.SetPowerUp(true)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = gorf24.B("01010101")
	c.SetPowerUp(false)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagZero_ReturnsFalse(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111101"))
	expected := false

	actual := c.IsPowerUp()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000010"))
	expected := true

	actual := c.IsPowerUp()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To16bits_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111011"))
	expected := gorf24.B("11111111")

	c.SetCrcLength(gorf24.CRC_16BIT)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("01010101"))

	expected := gorf24.B("01010001")
	c.SetCrcLength(gorf24.CRC_8BIT)
	actual := c.Byte()

	if actual != expected  {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = gorf24.B("01010101")
	c.SetCrcLength(gorf24.CRC_16BIT)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To8bits_SetsFlagToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000100"))
	expected := gorf24.B("00000000")

	c.SetCrcLength(gorf24.CRC_8BIT)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_8bits_Returns8bits(t *testing.T) {
	c := gorf24.NewConfigReg(0)
	expected := gorf24.CrcLength(gorf24.CRC_8BIT)
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_16bits_Returns16bits(t *testing.T) {
	c := gorf24.NewConfigReg(0)
	expected := gorf24.CrcLength(gorf24.CRC_16BIT)
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_True_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11110111"))
	expected := gorf24.B("11111111")

	c.SetCrcEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_False_SetsFlagToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00001000"))
	expected := gorf24.B("00000000")

	c.SetCrcEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("01010101"))

	expected := gorf24.B("01011101")
	c.SetCrcEnabled(true)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = gorf24.B("01010101")
	c.SetCrcEnabled(false)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b'but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00010000"))
	expected := gorf24.B("00000000")

	c.SetMaxRtInterruptEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11101111"))
	expected := gorf24.B("11111111")

	c.SetMaxRtInterruptEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))

	expected := gorf24.B("10111010")
	c.SetMaxRtInterruptEnabled(false)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = gorf24.B("10101010")
	c.SetMaxRtInterruptEnabled(true)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00100000"))
	expected := gorf24.B("00000000")

	c.SetTxDsInterruptEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11011111"))
	expected := gorf24.B("11111111")

	c.SetTxDsInterruptEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("01010101"))

	expected := gorf24.B("01110101")
	c.SetTxDsInterruptEnabled(false)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = gorf24.B("01010101")
	c.SetTxDsInterruptEnabled(true)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}








