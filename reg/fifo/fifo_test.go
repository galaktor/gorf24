package fifo

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNewFifoStatus_RegisterAddress_IsSTATUS(t *testing.T) {
	f := New(0)
	expected := addr.FIFO_STATUS

	actual := f.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_FullBitOne_ReturnsFULL(t *testing.T) {
	f := New(util.B("00000010"))
	expected := FULL

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyBitOne_ReturnsEMPTY(t *testing.T) {
	f := New(util.B("00000001"))
	expected := EMPTY

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyAndFullBitsZero_ReturnsPARTIAL(t *testing.T) {
	f := New(util.B("11111100"))
	expected := PARTIAL

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyAndFullBitsOne_ReturnsINVALID(t *testing.T) {
	f := New(util.B("00000011"))
	expected := INVALID

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_FullBitOne_ReturnsFULL(t *testing.T) {
	f := New(util.B("00100000"))
	expected := FULL

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyBitOne_ReturnsEMPTY(t *testing.T) {
	f := New(util.B("00010000"))
	expected := EMPTY

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyAndFullBitsZero_ReturnsPARTIAL(t *testing.T) {
	f := New(util.B("11001111"))
	expected := PARTIAL

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyAndFullBitsOne_ReturnsINVALID(t *testing.T) {
	f := New(util.B("00110000"))
	expected := INVALID

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestIsTxPayloadReuseEnabled_BitZero_ReturnsFalse(t *testing.T) {
	f := New(util.B("10111111"))
	expected := false

	actual := f.IsTxPayloadReuseEnabled()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fifostatus '%v'", expected, actual, f)
	}
}

func TestIsTxPayloadReuseEnabled_BitOne_ReturnsTrue(t *testing.T) {
	f := New(util.B("01000000"))
	expected := true

	actual := f.IsTxPayloadReuseEnabled()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fifostatus '%v'", expected, actual, f)
	}
}





