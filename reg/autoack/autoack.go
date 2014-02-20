/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package autoack

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* EN_AA
   Enable ‘Auto Acknowledgment’ Function Disable
   this functionality to be compatible with nRF2401
   bits 7:6 reserved */
type AA struct {
	reg.R
}

const RES_MASK byte = 0x3F // 00111111

func New(flags byte) *AA {
	masked := flags & RES_MASK
	return &AA{reg.New(addr.EN_AA, masked)}
}

// TODO: if pipe const value was byte, could just do
// "return a.flags & p == p"

/* ENAA_Px
   Enable auto acknowledgement data pipe 'p' */
func (a *AA) Set(p pipe.P, enabled bool) {
	if enabled {
		a.enable(p)
	} else {
		a.disable(p)
	}
}
func (a *AA) enable(p pipe.P) {
	switch p {
	case pipe.P0:
		a.R.Set(a.Byte() | 1)
	case pipe.P1:
		a.R.Set(a.Byte() | 2)
	case pipe.P2:
		a.R.Set(a.Byte() | 4)
	case pipe.P3:
		a.R.Set(a.Byte() | 8)
	case pipe.P4:
		a.R.Set(a.Byte() | 16)
	case pipe.P5:
		a.R.Set(a.Byte() | 32)
	}
}
func (a *AA) disable(p pipe.P) {
	switch p {
	case pipe.P0:
		a.R.Set(a.Byte() & 0xFE)
	case pipe.P1:
		a.R.Set(a.Byte() & 0xFD)
	case pipe.P2:
		a.R.Set(a.Byte() & 0xFB)
	case pipe.P3:
		a.R.Set(a.Byte() & 0xF7)
	case pipe.P4:
		a.R.Set(a.Byte() & 0xEF)
	case pipe.P5:
		a.R.Set(a.Byte() & 0xDF)
	}
}
func (a *AA) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0:
		result = a.Byte()&1 == 1
	case pipe.P1:
		result = a.Byte()&2 == 2
	case pipe.P2:
		result = a.Byte()&4 == 4
	case pipe.P3:
		result = a.Byte()&8 == 8
	case pipe.P4:
		result = a.Byte()&16 == 16
	case pipe.P5:
		result = a.Byte()&32 == 32
	}

	return result
}
