package reg

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

/* EN_AA
   Enable ‘Auto Acknowledgment’ Function Disable
   this functionality to be compatible with nRF2401
   bits 7:6 reserved */
type AutoAck struct {
	R
}

func NewAutoAck(flags byte) *AutoAck {
	return &AutoAck{R{a: addr.EN_AA, flags: flags}}
}

// TODO: if pipe const value was byte, could just do
// "return a.flags & p == p"

/* ENAA_Px 
   Enable auto acknowledgement data pipe 'p' */
func (a *AutoAck) Set(p pipe.P, enabled bool) {
	if enabled {
		a.enable(p)
	} else {
		a.disable(p)
	}
}
func (a *AutoAck) enable(p pipe.P) {
	switch p {
	case pipe.P0:
		a.flags |= 1
	case pipe.P1:
		a.flags |= 2
	case pipe.P2:
		a.flags |= 4
	case pipe.P3:
		a.flags |= 8
	case pipe.P4:
		a.flags |= 16
	case pipe.P5:
		a.flags |= 32
	}
}
func (a *AutoAck) disable(p pipe.P) {
	switch p {
	case pipe.P0:
		a.flags &= 0xFE
	case pipe.P1:
		a.flags &= 0xFD
	case pipe.P2:
		a.flags &= 0xFB
	case pipe.P3:
		a.flags &= 0xF7
	case pipe.P4:
		a.flags &= 0xEF
	case pipe.P5:
		a.flags &= 0xDF
	}
}
func (a *AutoAck) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0:
		result = a.flags&1 == 1
	case pipe.P1:
		result = a.flags&2 == 2
	case pipe.P2:
		result = a.flags&4 == 4
	case pipe.P3:
		result = a.flags&8 == 8
	case pipe.P4:
		result = a.flags&16 == 16
	case pipe.P5:
		result = a.flags&32 == 32
	}

	return result
}
