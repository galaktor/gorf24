/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"testing"
)

func TestNew_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := A(0x0000000000)
	
	actual := New(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := A(0xFFFFFFFFFF)
	
	actual := New(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := A(0xAAAAAAAAAA)
	
	actual := New(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNew_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := A(0xAAAAAAAAAA)

	actual := New(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := New(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















