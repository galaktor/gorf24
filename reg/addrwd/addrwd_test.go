package addrwd

import (
	"testing"

	"github.com/galaktor/gorf24/util"
)

func TestSet_3Bytes_FlipsRightBit(t *testing.T) {
	w := NewAddrWidths(util.B("00000000"))
	expected := util.B("00000001")

	w.Set(AW_3BYTES)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_4Bytes_FlipsRightBit(t *testing.T) {
	w := NewAddrWidths(util.B("00000000"))
	expected := util.B("00000010")

	w.Set(AW_4BYTES)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_5Bytes_FlipsRightBit(t *testing.T) {
	w := NewAddrWidths(util.B("00000000"))
	expected := util.B("00000011")

	w.Set(AW_5BYTES)

	actual := w.Byte()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_GreaterThree_ReturnsError(t *testing.T) {
	w := NewAddrWidths(util.B("00000000"))
	expected := "value outside of legal range: 4. only values 1 - 3 allowed."
	
	err := w.Set(AddrWidth(4))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with addrwidths '%v'", expected, actual, w)
	}
}

func TestSet_Zero_ReturnsError(t *testing.T) {
	w := NewAddrWidths(util.B("00000000"))
	expected := "value outside of legal range: 0. only values 1 - 3 allowed."
	
	err := w.Set(AddrWidth(0))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueOne_Returns3Bytes(t *testing.T) {
	w := NewAddrWidths(util.B("00000001"))
	expected := AW_3BYTES

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueTwo_Returns4Bytes(t *testing.T) {
	w := NewAddrWidths(util.B("00000010"))
	expected := AW_4BYTES

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}

func TestGet_ValueThree_Returns5Bytes(t *testing.T) {
	w := NewAddrWidths(util.B("00000011"))
	expected := AW_5BYTES

	actual := w.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with addrwidths '%v'", expected, actual, w)
	}
}
