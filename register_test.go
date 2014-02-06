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

func TestConfigReg_SetPrimaryReceiver_RelevantFlagOne_LeavesOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
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

func TestConfigReg_SetPrimaryTransmitter_FlagAlreadyZero_LeavesItZero(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("11111110"))
	expected := gorf24.B("11111110")

	c.SetPrimaryTransmitter()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}




















