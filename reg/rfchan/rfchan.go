/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
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

func New() *R {
	return &R{reg.New(addr.RF_CH, 0x7F)} // 01111111
}

func NewWith(flags byte) *R {
	r := New()
	r.R.Set(flags)
	return r
}

/* RF_CH (bits 6:0)
   Sets the frequency channel nRF24L01+ operates on */
func (c *R) Get() uint8 {
	return c.R.Get() & 0x7F // ignore bit 7
}
func (c *R) Set(ch uint8) error {
	if ch > 127 {
		return errors.New(fmt.Sprintf("value outside of legal range: %v. Only values 1 - 127 allowed.", ch))
	}

	c.R.Set(ch & 0x7F)
	return nil
}










