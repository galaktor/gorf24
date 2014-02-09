package cmd

import (
	"testing"

	"github.com/galaktor/gorf24/util"
)

func TestByte_Zero_ReturnsZero(t *testing.T) {
	c := cmd.C(util.B("00000000"))
	expected := util.B("00000000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_AllOnes_ReturnsSame(t *testing.T) {
	c := cmd.C(util.B("11111111"))
	expected := util.B("11111111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_LSByteZero_ReturnsSame(t *testing.T) {
	c := cmd.C(util.B("11110000"))
	expected := util.B("11110000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_MSByteZero_ReturnsSame(t *testing.T) {
	c := cmd.C(util.B("00001111"))
	expected := util.B("00001111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestCMD_R_REGISTER_ZeroBits_ReturnsZero(t *testing.T) {
	r := gorf24.Register(util.B("00000000"))
	expected := util.B("00000000")

	result := gorf24.CMD_R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_R_REGISTER_AllOnes_ReturnsLastFiveBits(t *testing.T) {
	r := gorf24.Register(util.B("11111111"))
	expected := util.B("00011111")

	result := gorf24.CMD_R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_W_REGISTER_ZeroBits_SetsWriteRegFlag(t *testing.T) {
	r := gorf24.Register(util.B("00000000"))
	expected := util.B("00100000")

	result := gorf24.CMD_W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_W_REGISTER_AllOnes_MasksTwoMSBits(t *testing.T) {
	r := gorf24.Register(util.B("11111111"))
	expected := util.B("00111111")

	result := gorf24.CMD_W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeZero_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(0)
	expected := util.B("10101000")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeOne_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(1)
	expected := util.B("10101001")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeTwo_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(2)
	expected := util.B("10101010")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeThree_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(3)
	expected := util.B("10101011")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeFour_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(4)
	expected := util.B("10101100")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeFive_MaskedCorrectly(t *testing.T) {
	p := gorf24.Pipe(5)
	expected := util.B("10101101")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeHigherThanFive_ThreeBits_MasksCorrectlyAlthoughInvalid(t *testing.T) {
	p := gorf24.Pipe(6)
	expected := util.B("10101110")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestCMD_W_ACK_PAYLOAD_PipeHigherThanFive_OverThreeBits_TruncatesAndMasksCorrectlyAlthoughInvalid(t *testing.T) {
	p := gorf24.Pipe(10)
	expected := util.B("10101010")
	
	result := gorf24.CMD_W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}
