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

func New(flags byte) *E {
	return &E{reg.New(addr.EN_RXADDR, flags)}
}

/* ERX_Px
   Enable data pipe 'p' */
func (e *E) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0:
		result = e.Byte()&1 == 1
	case pipe.P1:
		result = e.Byte()&2 == 2
	case pipe.P2:
		result = e.Byte()&4 == 4
	case pipe.P3:
		result = e.Byte()&8 == 8
	case pipe.P4:
		result = e.Byte()&16 == 16
	case pipe.P5:
		result = e.Byte()&32 == 32
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
		e.R.Set(e.Byte() | 1)
	case pipe.P1:
		e.R.Set(e.Byte() | 2)
	case pipe.P2:
		e.R.Set(e.Byte() | 4)
	case pipe.P3:
		e.R.Set(e.Byte() | 8)
	case pipe.P4:
		e.R.Set(e.Byte() | 16)
	case pipe.P5:
		e.R.Set(e.Byte() | 32)
	}
}
func (e *E) disable(p pipe.P) {
	switch p {
	case pipe.P0:
		e.R.Set(e.Byte() & 0xFE)
	case pipe.P1:
		e.R.Set(e.Byte() & 0xFD)
	case pipe.P2:
		e.R.Set(e.Byte() & 0xFB)
	case pipe.P3:
		e.R.Set(e.Byte() & 0xF7)
	case pipe.P4:
		e.R.Set(e.Byte() & 0xEF)
	case pipe.P5:
		e.R.Set(e.Byte() & 0xDF)
	}
}
