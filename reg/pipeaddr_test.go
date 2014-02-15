package reg

import (
	"testing"
)

func TestNewPipeAddress_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := PipeAddress(0x0000000000)
	
	actual := NewPipeAddress(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewPipeAddress_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := PipeAddress(0xFFFFFFFFFF)
	
	actual := NewPipeAddress(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewPipeAddress_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := PipeAddress(0xAAAAAAAAAA)
	
	actual := NewPipeAddress(0xAAAAAAAAAA)


	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewPipeAddress_OverflowFiveBytes_TruncatesToFiveBytes(t *testing.T) {
	expected := PipeAddress(0xAAAAAAAAAA)

	actual := NewPipeAddress(0xFFFFFFAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestByte_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := NewPipeAddress(expected)

	actual := a.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}



















