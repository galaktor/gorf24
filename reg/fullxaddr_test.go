package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
)

func TestGet_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := XAddress(0)
	a := newFullXAddress(addr.A(0), 0)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_NonZero_ReturnsRightBits(t *testing.T) {
	expected := XAddress(0xAFAFAFAFAF)
	a := newFullXAddress(addr.A(0), 0xAFAFAFAFAF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_StoresRightBits(t *testing.T) {
	expected := XAddress(0xAFAFAFAFAF)
	a := newFullXAddress(addr.A(0), XAddress(0))

	a.Set(0xAFAFAFAFAF)

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}
