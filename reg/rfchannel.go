package reg

import (
	"errors"
	"fmt"
	
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
func (c *RfChannel) Set(ch uint8) error {
	if ch > 127 {
		return errors.New(fmt.Sprintf("value outside of legal range: %v. Only values 1 - 127 allowed.", ch))
	}
	
	c.flags = ch & 0x7F
	return nil
}










