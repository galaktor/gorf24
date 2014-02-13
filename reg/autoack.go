package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type AutoAck struct {
	R
}

func NewAutoAck(flags byte) *AutoAck {
	return &AutoAck{R{a: addr.EN_AA, flags: flags}}
}

/* ENAA_Px */
func (a *AutoAck) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0: result = a.flags & 1 == 1
	case pipe.P1: result = a.flags & 2 == 2
	case pipe.P2: result = a.flags & 4 == 4
	case pipe.P3: result = a.flags & 8 == 8
	case pipe.P4: result = a.flags & 16 == 16
	case pipe.P5: result = a.flags & 32 == 32
	}

	return result
}