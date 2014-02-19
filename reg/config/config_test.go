package config

import (
	"testing"

	"github.com/galaktor/gorf24/util"
)

func TestSetPrimaryReceiver_RelevantFlagZero_SetsFlagToOne(t *testing.T) {
	c := New(util.B("00000000"))
	expected := util.B("00000001")

	c.SetPrimaryReceiver()

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryReceiver_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("10101010"))
	expected := util.B("10101011")

	c.SetPrimaryReceiver()

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagZero_ReturnsFalse(t *testing.T) {
	c := New(util.B("11111110"))
	expected := false

	actual := c.IsPrimaryReceiver()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryReceiver_FlagOne_ReturnsTrue(t *testing.T) {
	c := New(util.B("00000001"))
	expected := true

	actual := c.IsPrimaryReceiver()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_SetsBitToZero(t *testing.T) {
	c := New(util.B("00000001"))
	expected := util.B("00000000")

	c.SetPrimaryTransmitter()

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPrimaryTransmitter_DoesntFlipOtherBits(t *testing.T) {
	c := New(util.B("10101010"))
	expected := util.B("10101010")

	c.SetPrimaryTransmitter()

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagZero_ReturnsTrue(t *testing.T) {
	c := New(util.B("11111110"))
	expected := true

	actual := c.IsPrimaryTransmitter()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPrimaryTransmitter_FlagOne_ReturnsFalse(t *testing.T) {
	c := New(util.B("00000001"))
	expected := false

	actual := c.IsPrimaryTransmitter()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_True_SetsFlagToOne(t *testing.T) {
	c := New(util.B("11111101"))
	expected := util.B("11111111")

	c.SetPowerUp(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_False_SetsFlagToZero(t *testing.T) {
	c := New(util.B("00000010"))
	expected := util.B("00000000")

	c.SetPowerUp(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetPowerUp_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("01010101"))

	expected := util.B("01010111")
	c.SetPowerUp(true)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetPowerUp(false)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagZero_ReturnsFalse(t *testing.T) {
	c := New(util.B("11111101"))
	expected := false

	actual := c.IsPowerUp()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsPowerUp_FlagOne_ReturnsTrue(t *testing.T) {
	c := New(util.B("00000010"))
	expected := true

	actual := c.IsPowerUp()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To16bits_SetsFlagToOne(t *testing.T) {
	c := New(util.B("11111011"))
	expected := util.B("11111111")

	c.SetCrcLength(CRC_16BIT)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("01010101"))

	expected := util.B("01010001")
	c.SetCrcLength(CRC_8BIT)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetCrcLength(CRC_16BIT)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcLength_To8bits_SetsFlagToZero(t *testing.T) {
	c := New(util.B("00000100"))
	expected := util.B("00000000")

	c.SetCrcLength(CRC_8BIT)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_8bits_Returns8bits(t *testing.T) {
	c := New(0)
	expected := CRC_8BIT
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestGetCrcLength_16bits_Returns16bits(t *testing.T) {
	c := New(0)
	expected := CRC_16BIT
	c.SetCrcLength(expected)

	actual := c.GetCrcLength()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_True_SetsFlagToOne(t *testing.T) {
	c := New(util.B("11110111"))
	expected := util.B("11111111")

	c.SetCrcEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_False_SetsFlagToZero(t *testing.T) {
	c := New(util.B("00001000"))
	expected := util.B("00000000")

	c.SetCrcEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetCrcEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("01010101"))

	expected := util.B("01011101")
	c.SetCrcEnabled(true)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetCrcEnabled(false)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b'but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsCrcEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := New(util.B("11110111"))
	expected := false

	actual := c.IsCrcEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsCrcEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := New(util.B("00001000"))
	expected := true

	actual := c.IsCrcEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := New(util.B("00010000"))
	expected := util.B("00000000")

	c.SetMaxRtInterruptEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := New(util.B("11101111"))
	expected := util.B("11111111")

	c.SetMaxRtInterruptEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetMaxRtInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("10101010"))

	expected := util.B("10111010")
	c.SetMaxRtInterruptEnabled(false)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("10101010")
	c.SetMaxRtInterruptEnabled(true)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsMaxRtInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := New(util.B("00010000"))
	expected := false

	actual := c.IsMaxRtInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsMaxRtInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := New(util.B("11101111"))
	expected := true

	actual := c.IsMaxRtInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := New(util.B("00100000"))
	expected := util.B("00000000")

	c.SetTxDsInterruptEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := New(util.B("11011111"))
	expected := util.B("11111111")

	c.SetTxDsInterruptEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetTxDsInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("01010101"))

	expected := util.B("01110101")
	c.SetTxDsInterruptEnabled(false)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("01010101")
	c.SetTxDsInterruptEnabled(true)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsTxDsInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := New(util.B("00100000"))
	expected := false

	actual := c.IsTxDsInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsTxDsInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := New(util.B("11011111"))
	expected := true

	actual := c.IsTxDsInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_True_SetsBitToZero(t *testing.T) {
	c := New(util.B("01000000"))
	expected := util.B("00000000")

	c.SetRxDrInterruptEnabled(true)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_False_SetsBitToOne(t *testing.T) {
	c := New(util.B("10111111"))
	expected := util.B("11111111")

	c.SetRxDrInterruptEnabled(false)

	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestSetRxDrInterruptEnabled_DoesNotFlipOtherBits(t *testing.T) {
	c := New(util.B("10101010"))

	expected := util.B("11101010")
	c.SetRxDrInterruptEnabled(false)
	actual := c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}

	expected = util.B("10101010")
	c.SetRxDrInterruptEnabled(true)
	actual = c.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestIsRxDrInterruptEnabled_Disabled_ReturnsFalse(t *testing.T) {
	c := New(util.B("01000000"))
	expected := false

	actual := c.IsRxDrInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestIsRxDrInterruptEnabled_Enabled_ReturnsTrue(t *testing.T) {
	c := New(util.B("10111111"))
	expected := true

	actual := c.IsRxDrInterruptEnabled()

	if actual != expected {
		t.Errorf("expected '%v'but found '%v' with configreg '%v'", expected, actual, c)
	}
}