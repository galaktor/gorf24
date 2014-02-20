/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package addrw

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

type Width byte

const (
	BYTES_3 Width = 0x1
	BYTES_4 Width = 0x2
	BYTES_5 Width = 0x3
)

/* SETUP_AW
   Setup of Address Widths
   (common for all data pipes)
   can only use two LSBits; 
   bits 7:2 reserved, must be 000000 */
type AW struct {
	reg.R
}

const RES_MASK byte = 0x03 // 00000011

func New(flags byte) *AW {
	masked := flags & RES_MASK
	return &AW{reg.New(addr.SETUP_AW, masked)}
}

/* AW (bits 1:0)
   RX/TX Address field width
   '00' - Illegal
   '01' - 3 bytes
   '10' - 4 bytes
   '11' – 5 bytes
   LSByte is used if address width is below 5 bytes */
func (a *AW) Set(w Width) error {
	if w == 0 || w&0xFC > 0 {
		return errors.New(fmt.Sprintf("value outside of legal range: %v. only values 1 - 3 allowed.", w))
	}

	a.R.Set(byte(w))
	return nil
}
func (a *AW) Get() Width {
	return Width(a.Byte())
}
