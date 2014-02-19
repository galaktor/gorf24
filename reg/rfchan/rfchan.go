package rfchan

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* RF_CH
   RF Channel
   bit 7 reserved */
type R struct {
	reg.R
}

func New(flags byte) *R {
	return &R{reg.New(addr.RF_CH, flags)}
}

/* RF_CH (bits 6:0)
   Sets the frequency channel nRF24L01+ operates on */
func (c *R) Get() uint8 {
	return c.Byte() & 0x7F // ignore bit 7
}
func (c *R) Set(ch uint8) error {
	if ch > 127 {
		return errors.New(fmt.Sprintf("value outside of legal range: %v. Only values 1 - 127 allowed.", ch))
	}

	c.R.Set(ch & 0x7F)
	return nil
}










