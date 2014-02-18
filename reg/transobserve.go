package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* OBSERVE_TX
   Transmit observe register */
type TransObserve struct {
	R
}

func NewTransObserve(flags byte) *TransObserve {
	return &TransObserve{R{a: addr.OBSERVE_TX, flags: flags}}
}

/* ARC_CNT (bits 3:0)
   Count retransmitted packets. The counter is reset
   when transmission of a new packet starts. */
func (o *TransObserve) RetransmittedPacketCount() uint8 {
	// mask out 4 MSbits
	return o.flags & 0x0F
}

/* PLOS_CNT (bits 7:4)
   Count lost packets. The counter is overflow pro-
   tected to 15, and discontinues at max until reset.
   The counter is reset by writing to RF_CH. */
func (o *TransObserve) LostPacketCount() uint8 {
	// dump 4 LSbits
	return (o.flags & 0xF0) >> 4
}
