package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
)

func someFullXAddr(flags uint64) *FullXAddress {
	return newFullXAddress(addr.A(0), NewXAddress(flags))
}

func TestByte_ParentMSBytesZero_FirstFourBytesZero(t *testing.T) {
	expected := XAddress(0x00000000FF)
	root := someFullXAddr(0x0000000000)
	a := newPartialXAddress(addr.A(0), root, 0xFF)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSBytesOnes_FirstFourBytesOnes(t *testing.T) {
	expected := XAddress(0xFFFFFFFFAA)
	root := someFullXAddr(0xFFFFFFFFFF)
	a := newPartialXAddress(addr.A(0), root, 0xAA)

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}

func TestByte_ParentMSByteChanges_ByteMsbChangesWithParent(t *testing.T) {
	expected := XAddress(0xFFFFFFFFA1)
	root := someFullXAddr(0x0000000000)
	a := newPartialXAddress(addr.A(0), root, 0xA1)
	
	root.Set(NewXAddress(0xFFFFFFFFFF))
	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with partaddr '%v'", expected, actual, a)
	}
}
