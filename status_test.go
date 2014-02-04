package gorf24

import (
	"testing"
	"strconv"

	"."
)

func b(bits string) byte {
	i,_ := strconv.ParseUint(bits, 2, 8)
	return byte(i)
}

func TestTxFull_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("11111110"))
	expected := false

	result := s.TxFull()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxFull_RelevantBitsOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status(b("00000001"))
	expected := true

	result := s.TxFull()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxPipeNumber_RelevantBitsZero_ReturnsZero(t *testing.T) {
	s := gorf24.Status(b("11110001"))
	expected := uint8(0)

	result := s.RxPipeNumber()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxPipeNumber_RelevantBitsOne_ReturnsSeven(t *testing.T) {
	s := gorf24.Status(b("00001110"))
	expected := uint8(7)
	
	result := s.RxPipeNumber()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxPipeNumberUsed_RelevantBitsFive_ReturnsTrue(t *testing.T) {
	s := gorf24.Status(b("00001010"))
	expected := true

	result := s.RxPipeNumberUsed()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxPipeNumberUsed_RelevantBitsSix_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("00001100"))
	expected := false

	result := s.RxPipeNumberUsed()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxFifoEmpty_RelevantBitsSix_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("00001100"))
	expected := false

	result := s.RxFifoEmpty()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxFifoEmpty_RelevantBitsSeven_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("00001110"))
	expected := true

	result := s.RxFifoEmpty()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestMaxTxRetransmits_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("11110111"))
	expected := false

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestMaxTxRetransmits_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status(b("00001000"))
	expected := true

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxDataSent_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("11101111"))
	expected := false

	result := s.TxDataSent()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestTxDataSent_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status(b("00010000"))
	expected := true

	result := s.TxDataSent()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxDataReady_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status(b("11011111"))
	expected := false

	result := s.RxDataReady()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}

func TestRxDataReady_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status(b("00100000"))
	expected := true

	result := s.RxDataReady()

	if result != expected {
		t.Errorf("expected '%v' but found '%v' with status '%v'", expected, result, s)
	}
}
