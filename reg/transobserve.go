package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type TransObserve struct {
	R
}

func NewTransObserve(flags byte) *TransObserve {
	return &TransObserve{R{a: addr.OBSERVE_TX, flags: flags}}
}

/* ARC_CNT (bits 3:0) */
func (o *TransObserve) RetransmittedPacketCount() uint8 {
	return o.flags & 0x0F
}