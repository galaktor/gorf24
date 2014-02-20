/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package dynpd

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* DYNPD
   Enable dynamic payload length
   bits 7:6 reserved */
type DP struct {
	reg.R
}

const RES_MASK byte = 0x3F // 00111111

func New(flags byte) *DP {
	masked := flags & RES_MASK
	return &DP{reg.New(addr.DYNPD, masked)}
}

/* DPL_Px
   Enable dynamic payload length data pipe 'p'.
   (Requires EN_DPL and ENAA_Px) */
func (d *DP) IsEnabled(p pipe.P) (enabled bool) {
	switch p {
	case pipe.P0:
		enabled = d.Byte()&0x01 == 0x01
	case pipe.P1:
		enabled = d.Byte()&0x02 == 0x02
	case pipe.P2:
		enabled = d.Byte()&0x04 == 0x04
	case pipe.P3:
		enabled = d.Byte()&0x08 == 0x08
	case pipe.P4:
		enabled = d.Byte()&0x10 == 0x10
	case pipe.P5:
		enabled = d.Byte()&0x20 == 0x20
	}

	return
}
func (d *DP) Set(p pipe.P, enabled bool) error {
	if enabled {
		return d.enable(p)
	} else {
		return d.disable(p)
	}
}
func (d *DP) enable(p pipe.P) (err error) {
	err = nil

	switch p {
	case pipe.P0:
		d.R.Set(d.Byte() | 0x01)
	case pipe.P1:
		d.R.Set(d.Byte() | 0x02)
	case pipe.P2:
		d.R.Set(d.Byte() | 0x04)
	case pipe.P3:
		d.R.Set(d.Byte() | 0x08)
	case pipe.P4:
		d.R.Set(d.Byte() | 0x10)
	case pipe.P5:
		d.R.Set(d.Byte() | 0x20)
	default:
		err = errors.New(fmt.Sprintf("invalid pipe: %v", p))
	}

	return
}
func (d *DP) disable(p pipe.P) (err error) {
	err = nil

	switch p {
	case pipe.P0:
		d.R.Set(d.Byte() & 0xFE)
	case pipe.P1:
		d.R.Set(d.Byte() & 0xFD)
	case pipe.P2:
		d.R.Set(d.Byte() & 0xFB)
	case pipe.P3:
		d.R.Set(d.Byte() & 0xF7)
	case pipe.P4:
		d.R.Set(d.Byte() & 0xEF)
	case pipe.P5:
		d.R.Set(d.Byte() & 0xDF)
	default:
		err = errors.New(fmt.Sprintf("invalid pipe: %v", p))
	}

	return
}
