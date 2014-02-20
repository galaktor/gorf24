/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rxpw

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

const MAX uint8 = 32

/* PX_PW_Px
   Number of bytes in RX payload in data pipe 0 (1 to
   32 bytes).
   0 Pipe not used
   1 = 1 byte
   ...
   32 = 32 bytes

   bits 7:6 reserved */
type W struct {
	reg.R
}

const RES_MASK byte = 0x3F // 00111111

func New(p pipe.P, flags byte) *W {
	masked := flags & RES_MASK
	return &W{reg.New(addr.RX_PW(p), masked)}
}

/* RX_PWx (bits 5:0) */
func (r *W) Set(w uint8) error {
	if w > MAX {
		return errors.New(fmt.Sprintf("value out of legal range: %v. allowed values from 0 - %v", w, MAX))
	}

	r.R.Set((r.Byte() & 0) | w)
	return nil
}
func (r *W) Get() uint8 {
	if r.Byte() > MAX {
		return MAX
	}

	return r.Byte()
}
