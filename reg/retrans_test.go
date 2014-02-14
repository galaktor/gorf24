package reg

import (
	"testing"

	"github.com/galaktor/gorf24/util"
)

func TestSetCount_Zero_FlipsRightBits(t *testing.T) {
	r := NewSetupRetrans(util.B("11111111"))
	expected := util.B("11110000")

	err := r.SetCount(0)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_One_FlipsRightBits(t *testing.T) {
	r := NewSetupRetrans(util.B("11110000"))
	expected := util.B("11110001")

	err := r.SetCount(1)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_One_ReturnsNoError(t *testing.T) {
	r := NewSetupRetrans(util.B("11110000"))

	err := r.SetCount(1)

	if err != nil {
		t.Errorf("unexpected error '%v' with retrans '%v'", err, r)
	}
}

func TestSetCount_Two_FlipsRightBits(t *testing.T) {
	r := NewSetupRetrans(util.B("11110000"))
	expected := util.B("11110010")

	err := r.SetCount(2)

	if err != nil  {
		t.Errorf("unexpected error '%v'", err)
	}

	actual := r.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_Fifteen_FlipsRightBits(t *testing.T) {
	r := NewSetupRetrans(util.B("11110000"))
	expected := util.B("11111111")

	r.SetCount(15)

	actual := r.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with retrans '%v'", expected, actual, r)
	}
}

func TestSetCount_Sixteen_ReturnsError(t *testing.T) {
	r := NewSetupRetrans(util.B("00000000"))
	expected := "value out of legal range: 16. Only values 0 - 15 allowed."
	
	err := r.SetCount(16)

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with retrans '%v'", expected, actual, r)
	}
}


















