/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package txobs

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* OBSERVE_TX
   Transmit observe register */
type O struct {
	reg.R
}

func New(flags byte) *O {
	return &O{reg.New(addr.OBSERVE_TX, flags)}
}

/* ARC_CNT (bits 3:0)
   Count retransmitted packets. The counter is reset
   when transmission of a new packet starts. */
func (o *O) RetransmittedPacketCount() uint8 {
	// mask out 4 MSbits
	return o.Byte() & 0x0F
}

/* PLOS_CNT (bits 7:4)
   Count lost packets. The counter is overflow pro-
   tected to 15, and discontinues at max until reset.
   The counter is reset by writing to RF_CH. */
func (o *O) LostPacketCount() uint8 {
	// dump 4 LSbits
	return (o.Byte() & 0xF0) >> 4
}
