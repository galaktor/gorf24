package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNewFifoStatus_RegisterAddress_IsFIFO_STATUS(t *testing.T) {
	f := NewFifoStatus(0)
	expected := addr.FIFO_STATUS

	actual := f.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_FullBitOne_ReturnsFIFO_FULL(t *testing.T) {
	f := NewFifoStatus(util.B("00000010"))
	expected := FIFO_FULL

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyBitOne_ReturnsFIFO_EMPTY(t *testing.T) {
	f := NewFifoStatus(util.B("00000001"))
	expected := FIFO_EMPTY

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyAndFullBitsZero_ReturnsFIFO_PARTIAL(t *testing.T) {
	f := NewFifoStatus(util.B("11111100"))
	expected := FIFO_PARTIAL

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestRx_EmptyAndFullBitsOne_ReturnsFIFO_INVALID(t *testing.T) {
	f := NewFifoStatus(util.B("00000011"))
	expected := FIFO_INVALID

	actual := f.Rx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_FullBitOne_ReturnsFIFO_FULL(t *testing.T) {
	f := NewFifoStatus(util.B("00100000"))
	expected := FIFO_FULL

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyBitOne_ReturnsFIFO_EMPTY(t *testing.T) {
	f := NewFifoStatus(util.B("00010000"))
	expected := FIFO_EMPTY

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyAndFullBitsZero_ReturnsFIFO_PARTIAL(t *testing.T) {
	f := NewFifoStatus(util.B("11001111"))
	expected := FIFO_PARTIAL

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestTx_EmptyAndFullBitsOne_ReturnsFIFO_INVALID(t *testing.T) {
	f := NewFifoStatus(util.B("00110000"))
	expected := FIFO_INVALID

	actual := f.Tx()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}

func TestIsTxPayloadReuseEnabled_BitZero_ReturnsFalse(t *testing.T) {
	f := NewFifoStatus(util.B("10111111"))
	expected := false

	actual := f.IsTxPayloadReuseEnabled()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fifostatus '%v'", expected, actual, f)
	}
}

func TestIsTxPayloadReuseEnabled_BitOne_ReturnsTrue(t *testing.T) {
	f := NewFifoStatus(util.B("01000000"))
	expected := true

	actual := f.IsTxPayloadReuseEnabled()
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fifostatus '%v'", expected, actual, f)
	}
}





