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

// TODO: if pipe const value was byte, could just do
// "return a.flags & p == p"

/* ENAA_Px */
func (a *AutoAck) Enable(p pipe.P) {
	switch p {
	case pipe.P0: a.flags = a.flags | 1
	case pipe.P1: a.flags = a.flags | 2
	case pipe.P2: a.flags = a.flags | 4
	case pipe.P3: a.flags = a.flags | 8
	case pipe.P4: a.flags = a.flags | 16
	case pipe.P5: a.flags = a.flags | 32
	}
}
func (a *AutoAck) Disable(p pipe.P) {
	switch p {
	case pipe.P0: a.flags = a.flags & 0xFE
	case pipe.P1: a.flags = a.flags & 0xFD
	case pipe.P2: a.flags = a.flags & 0xFB
	case pipe.P3: a.flags = a.flags & 0xF7
	case pipe.P4: a.flags = a.flags & 0xEF
	case pipe.P5: a.flags = a.flags & 0xDF
	}
}
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
