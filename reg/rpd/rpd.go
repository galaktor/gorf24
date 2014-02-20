/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rpd

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* RPD
   Received Power Detector. This register is called
   CD (Carrier Detect) in the nRF24L01. The name is
   different in nRF24L01+ due to the different input
   power level threshold for this bit.
   bits 7:1 reserved */
type R struct {
	reg.R
}

const RES_MASK byte = 0x01

func New(flags byte) *R {
	masked := flags & RES_MASK
	return &R{reg.New(addr.RPD, masked)}
}

/* RPD (bit 0)
   triggers at received power levels above -64 dBm that
   are present in the RF channel you receive on. If the
   received power is less than -64 dBm, [sic] RDP = 0.*/
func (r *R) Triggered() bool {
	return r.Byte()&1 == 1
}
