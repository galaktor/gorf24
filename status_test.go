package gorf24

import (
	"testing"

	"."
)

const (
	b00000000 byte = 0x00
	b00000001 = 0x01
	b11111110 = ^0x01
	b11111111 = 0xFF
	
)

func TestTxFull_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b11111110}
	expected := false

	result := s.TxFull()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestTxFull_RelevantBitsOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status{b00000001}
	expected := true

	result := s.TxFull()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxPipeNumber_RelevantBitsZero_ReturnsZero(t *testing.T) {
	s := gorf24.Status{b11110001}
	expected = uint8(0)

	result := s.RxPipeNumber()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxPipeNumber_RelevantBitsOne_ReturnsSeven(t *testing.T) {
	s := gorf24.Status{b00001110}
	expected = uint8(7)
	
	result := s.RxPipeNumber()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxPipeNumberUsed_RelevantBitsFive_ReturnsTrue(t *testing.T) {
	s := gorf24.Status{b00001010}
	expected := true

	result := s.RxPipeNumberUsed()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxPipeNumberUsed_RelevantBitsSix_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b00001100}
	expected := false

	result := s.RxPipeNumberUsed()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxFifoEmpty_RelevantBitsSix_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b00001100}
	expected := false

	result := s.RxFifoEmpty()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxFifoEmpty_RelevantBitsSeven_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b00001110}
	expected := true

	result := s.RxFifoEmpty()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestMaxTxRetransmits_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b11110111}
	expected := false

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestMaxTxRetransmits_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status{b00001000}
	expected := true

	result := s.MaxTxRetransmits()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestTxDataSent_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b11101111}
	expected := false

	result := s.TxDataSent()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestTxDataSent_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status{b00010000}
	expected := true

	result := s.TxDataSent()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxDataReady_RelevantBitZero_ReturnsFalse(t *testing.T) {
	s := gorf24.Status{b11011111}
	expected := false

	result := s.RxDataReady()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}

func TestRxDataReady_RelevantBitOne_ReturnsTrue(t *testing.T) {
	s := gorf24.Status{b00100000}
	expected := true

	result := s.RxDataReady()

	if result != expected {
		t.Error("expected %v but found %v", expected, result)
	}
}










