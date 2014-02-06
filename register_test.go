package gorf24

import (
	"testing"

	"."
)

func TestConfigReg_SetPrimaryReader_RelevantFlagZero_SetsFlagToOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000000"))
	expected := gorf24.B("00000001")

	c.SetPrimaryReader()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPrimaryReader_RelevantFlagOne_LeavesOne(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("00000001"))
	expected := gorf24.B("00000001")

	c.SetPrimaryReader()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}

func TestConfigReg_SetPrimaryReader_SomeIrrelevantFlagsFlipped_ChangesOnlyRelevantBit(t *testing.T) {
	c := gorf24.NewConfigReg(gorf24.B("10101010"))
	expected := gorf24.B("10101011")

	c.SetPrimaryReader()
	
	actual := c.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with configreg '%v'", expected, actual, c)
	}
}




















