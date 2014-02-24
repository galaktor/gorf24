/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package enrxaddr

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* EN_RXADDR
   Enable RX Addresses
   bits 7:6 reserved */
type E struct {
	reg.R
}

func New() *E {
	return &E{reg.New(addr.EN_RXADDR, 0x3F)} // 00111111
}

func NewWith(flags byte) *E {
	e := New()
	e.R.Set(flags)
	return e
}

/* ERX_Px
   Enable data pipe 'p' */
func (e *E) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0:
		result = e.R.Get()&1 == 1
	case pipe.P1:
		result = e.R.Get()&2 == 2
	case pipe.P2:
		result = e.R.Get()&4 == 4
	case pipe.P3:
		result = e.R.Get()&8 == 8
	case pipe.P4:
		result = e.R.Get()&16 == 16
	case pipe.P5:
		result = e.R.Get()&32 == 32
	}

	return result
}
func (e *E) Set(p pipe.P, enabled bool) {
	if enabled {
		e.enable(p)
	} else {
		e.disable(p)
	}
}
func (e *E) enable(p pipe.P) {
	switch p {
	case pipe.P0:
		e.R.Set(e.R.Get() | 1)
	case pipe.P1:
		e.R.Set(e.R.Get() | 2)
	case pipe.P2:
		e.R.Set(e.R.Get() | 4)
	case pipe.P3:
		e.R.Set(e.R.Get() | 8)
	case pipe.P4:
		e.R.Set(e.R.Get() | 16)
	case pipe.P5:
		e.R.Set(e.R.Get() | 32)
	}
}
func (e *E) disable(p pipe.P) {
	switch p {
	case pipe.P0:
		e.R.Set(e.R.Get() & 0xFE)
	case pipe.P1:
		e.R.Set(e.R.Get() & 0xFD)
	case pipe.P2:
		e.R.Set(e.R.Get() & 0xFB)
	case pipe.P3:
		e.R.Set(e.R.Get() & 0xF7)
	case pipe.P4:
		e.R.Set(e.R.Get() & 0xEF)
	case pipe.P5:
		e.R.Set(e.R.Get() & 0xDF)
	}
}
