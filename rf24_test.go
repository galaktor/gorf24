/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gorf24

import (
	"testing"

	"github.com/galaktor/gorf24/gpio"
	"github.com/galaktor/gorf24/spi"
	"github.com/galaktor/gorf24/util"
)

func TestB_ZeroByte_ReturnsZero(t *testing.T) {
	bits := "00000000"
	expected := byte(0)

	result := util.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with bits '%s'", expected, result, bits)
	}
}

func TestB_AllOnes_Returns255(t *testing.T) {
	bits := "11111111"
	expected := byte(255)

	result := util.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}

func TestB_OverflowWithNineBits_AllOnes_Returns255(t *testing.T) {
	bits := "111111111"
	expected := byte(255)

	result := util.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}

func TestB_FortyTwo_ReturnsFortyTwo(t *testing.T) {
	bits := "101010"
	expected := byte(42)

	result := util.B(bits)

	if result != expected {
		t.Errorf("expected '%b' but found '%b' with command '%s'", expected, result, bits)
	}
}

func TestNew_ReadAllRegisters_SpiFullOfOnes_AllRegistersReadOnes(t *testing.T) {
	s := spi.NewFake(util.B("1111111"))
	ce := &gpio.Fake{}
	csn := &gpio.Fake{}

	r, err := New(s, ce, csn)

	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	r.ReadAllRegisters()

	// check config
	if r.config.Get() != util.B("1111111") {
		t.Errorf("wrong bits! spi: %v  config: %v", s, r.config)
	}
}



















