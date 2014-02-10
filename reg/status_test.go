package reg

import (
	"testing"

	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/pipe"
)

func TestTxFull_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("11111110"))
	expected := false

	result := s.TxFull()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxFull_RelevantBitsOne_ReturnsTrue(t *testing.T) {
	s := NewStatus(util.B("00000001"))
	expected := true

	result := s.TxFull()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxPipeNumber_RelevantBitsZero_ReturnsP0(t *testing.T) {
	s := NewStatus(util.B("11110001"))
	expected := pipe.P0

	result,err := s.RxPipeNumber()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with status '%b'", expected, result, s)
	}
}

func TestRxPipeNumber_RelevantBitsFive_ReturnsP5(t *testing.T) {
	s := NewStatus(util.B("00001010"))
	expected := pipe.P5
	
	result,err := s.RxPipeNumber()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with status '%b'", expected, result, s)
	}
}

func TestRxPipeNumber_RelevantBitsSix_ReturnsError(t *testing.T) {
	s := NewStatus(util.B("00001100"))
	expected := "Invalid pipe number: 6"
	
	_,err := s.RxPipeNumber()

	if err.Error() != expected {
		t.Errorf("expected error '%s' but found '%s'", expected, err.Error())
	}
}

func TestRxPipeNumber_RelevantBitsSeven_ReturnsError(t *testing.T) {
	s := NewStatus(util.B("00001110"))
	expected := "Rx FIFO is empty."
	
	_,err := s.RxPipeNumber()

	if err.Error() != expected {
		t.Errorf("expected error '%s' but found '%s'", expected, err.Error())
	}
}

func TestRxFifoEmpty_RelevantBitsSix_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("00001100"))
	expected := false

	result := s.RxFifoEmpty()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxFifoEmpty_RelevantBitsSeven_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("00001110"))
	expected := true

	result := s.RxFifoEmpty()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestMaxTxRetransmits_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("11110111"))
	expected := false

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestMaxTxRetransmits_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := NewStatus(util.B("00001000"))
	expected := true

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxDataSent_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("11101111"))
	expected := false

	result := s.TxDataSent()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxDataSent_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := NewStatus(util.B("00010000"))
	expected := true

	result := s.TxDataSent()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxDataReady_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := NewStatus(util.B("11011111"))
	expected := false

	result := s.RxDataReady()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxDataReady_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := NewStatus(util.B("00100000"))
	expected := true

	result := s.RxDataReady()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}
