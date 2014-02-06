package gorf24

import (
	"testing"

	"."
)

func TestConfigReg_SetPrimaryReceiver_RelevantFlagZero_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000000"))
	expected := gorf24.B("00000001")

	c.SetPrimaryReceiver()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPrimaryReceiver_SomeIrrelevantFlagsFlipped_ChangesOnlyRelevantBit(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101011")

	c.SetPrimaryReceiver()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPrimaryReceiver_FlagZero_ReturnsFalse(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111110"))
	expected := false

	actual := c.IsPrimaryReceiver()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPrimaryReceiver_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := true

	actual := c.IsPrimaryReceiver()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPrimaryTransmitter_SetsBitToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := gorf24.B("00000000")

	c.SetPrimaryTransmitter()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPrimaryTransmitter_SomeOtherBitsFlipped_DoesntChangeThose(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101010")

	c.SetPrimaryTransmitter()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPrimaryTransmitter_FlagZero_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111110"))
	expected := true

	actual := c.IsPrimaryTransmitter()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPrimaryTransmitter_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := false

	actual := c.IsPrimaryTransmitter()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPowerUp_FlagZero_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111101"))
	expected := gorf24.B("11111111")

	c.SetPowerUp()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPowerUp_OtherBitsSet_LeavesThoseAlone(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("01010101"))
	expected := gorf24.B("01010111")

	c.SetPowerUp()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPowerUp_FlagZero_ReturnsFalse(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111101"))
	expected := false

	actual := c.IsPowerUp()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPowerUp_FlagOne_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000010"))
	expected := true

	actual := c.IsPowerUp()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPowerDown_FlagOne_SetsFlagToZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000010"))
	expected := gorf24.B("00000000")

	c.SetPowerDown()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPowerDown_OtherBitsFlipped_LeavesThemAlone(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101000")

	c.SetPowerDown()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPowerDown_FlagZero_ReturnsTrue(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111101"))
	expected := true

	actual := c.IsPowerDown()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_IsPowerDown_FlagOne_ReturnsFalse(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000010"))
	expected := false

	actual := c.IsPowerDown()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with configreg '%v'", expected, actual, c)
	}
}




















