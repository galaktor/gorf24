package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

/* DYNPD
   Enable dynamic payload length
   bits 7:6 reserved */
type DynamicPayload struct {
	R
}

func NewDynamicPayload(flags byte) *DynamicPayload {
	return &DynamicPayload{R{a: addr.DYNPD, flags: flags}}
}

/* DPL_Px
   requires EN_DPL and ENAA_Px */
func (d *DynamicPayload) IsEnabled(p pipe.P) (enabled bool) {
	switch p {
	case pipe.P0:
		enabled = d.flags&0x01 == 0x01
	case pipe.P1:
		enabled = d.flags&0x02 == 0x02
	case pipe.P2:
		enabled = d.flags&0x04 == 0x04
	case pipe.P3:
		enabled = d.flags&0x08 == 0x08
	case pipe.P4:
		enabled = d.flags&0x10 == 0x10
	case pipe.P5:
		enabled = d.flags&0x20 == 0x20
	}

	return
}
func (d *DynamicPayload) Set(p pipe.P, enabled bool) error {
	if enabled {
		return d.enable(p)
	} else {
		return d.disable(p)
	}
}
func (d *DynamicPayload) enable(p pipe.P) (err error) {
	err = nil

	switch p {
	case pipe.P0:
		d.flags |= 0x01
	case pipe.P1:
		d.flags |= 0x02
	case pipe.P2:
		d.flags |= 0x04
	case pipe.P3:
		d.flags |= 0x08
	case pipe.P4:
		d.flags |= 0x10
	case pipe.P5:
		d.flags |= 0x20
	default:
		err = errors.New(fmt.Sprintf("invalid pipe: %v", p))
	}

	return
}
func (d *DynamicPayload) disable(p pipe.P) (err error) {
	err = nil

	switch p {
	case pipe.P0:
		d.flags &= 0xFE
	case pipe.P1:
		d.flags &= 0xFD
	case pipe.P2:
		d.flags &= 0xFB
	case pipe.P3:
		d.flags &= 0xF7
	case pipe.P4:
		d.flags &= 0xEF
	case pipe.P5:
		d.flags &= 0xDF
	default:
		err = errors.New(fmt.Sprintf("invalid pipe: %v", p))
	}

	return
}
