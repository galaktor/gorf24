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











