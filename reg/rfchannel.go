package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

// bit 7 always zero

type RfChannel struct {
	R
}

func NewRfChannel(flags byte) *RfChannel {
	return &RfChannel{R{a: addr.RF_CH, flags: flags}}
}



/* RF_CH */
func (c *RfChannel) Get() uint8 {
	return c.flags & 0x7F // ignore bit 7
}