/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package cmd

import (
	"testing"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
)

func TestByte_Zero_ReturnsZero(t *testing.T) {
	c := C(util.B("00000000"))
	expected := util.B("00000000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_AllOnes_ReturnsSame(t *testing.T) {
	c := C(util.B("11111111"))
	expected := util.B("11111111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_LSByteZero_ReturnsSame(t *testing.T) {
	c := C(util.B("11110000"))
	expected := util.B("11110000")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestByte_MSByteZero_ReturnsSame(t *testing.T) {
	c := C(util.B("00001111"))
	expected := util.B("00001111")

	result := c.Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%b'", expected, result, c)
	}
}

func TestR_REGISTER_AllZero_ReturnsZero(t *testing.T) {
	r := someReg(util.B("00000000"))
	expected := util.B("00000000")

	result := R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestR_REGISTER_AllOnes_ReturnsLastFiveBits(t *testing.T) {
	r := someReg(util.B("11111111"))
	expected := util.B("00011111")

	result := R_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestW_REGISTER_ZeroBits_SetsWriteRegFlag(t *testing.T) {
	r := someReg(util.B("00000000"))
	expected := util.B("00100000")

	result := W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestW_REGISTER_AllOnes_MasksTwoMSBits(t *testing.T) {
	r := someReg(util.B("11111111"))
	expected := util.B("00111111")

	result := W_REGISTER(r).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with register '%b'", expected, result, r)
	}
}

func TestW_ACK_PAYLOAD_PipeZero_MaskedCorrectly(t *testing.T) {
	p := pipe.P0
	expected := util.B("10101000")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeOne_MaskedCorrectly(t *testing.T) {
	p := pipe.P1
	expected := util.B("10101001")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeTwo_MaskedCorrectly(t *testing.T) {
	p := pipe.P2
	expected := util.B("10101010")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeThree_MaskedCorrectly(t *testing.T) {
	p := pipe.P3
	expected := util.B("10101011")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeFour_MaskedCorrectly(t *testing.T) {
	p := pipe.P4
	expected := util.B("10101100")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeFive_MaskedCorrectly(t *testing.T) {
	p := pipe.P5
	expected := util.B("10101101")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeHigherThanFive_ThreeBits_MasksCorrectlyAlthoughInvalid(t *testing.T) {
	p := pipe.P(6)
	expected := util.B("10101110")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestW_ACK_PAYLOAD_PipeHigherThanFive_OverThreeBits_TruncatesAndMasksCorrectlyAlthoughInvalid(t *testing.T) {
	p := pipe.P(10)
	expected := util.B("10101010")
	
	result := W_ACK_PAYLOAD(p).Byte()

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with pipe '%b'", expected, result, p)
	}
}

func TestNOP_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("11111111")
	
	actual := NOP.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestR_RX_PAYLOAD_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("01100001")
	
	actual := R_RX_PAYLOAD.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestW_TX_PAYLOAD_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("10100000")

	actual := W_TX_PAYLOAD.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestFLUSH_TX_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("11100001")

	actual := FLUSH_TX.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestFLUSH_RX_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("11100010")

	actual := FLUSH_RX.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestREUSE_TX_PL_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("11100011")

	actual := REUSE_TX_PL.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestACTIVATE_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("01010000")

	actual := ACTIVATE.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestR_RX_PL_WID_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("01100000")

	actual := R_RX_PL_WID.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestW_TX_PAYLOAD_NOACK_Always_ReturnsRightBits(t *testing.T) {
	expected := util.B("10110000")

	actual := W_TX_PAYLOAD_NOACK.Byte()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

/***** helper funcs *****/

func someReg(adr byte) reg.R {
	return reg.New(addr.A(adr), 0)
}














