package reg

import (
	"errors"
	"fmt"
	
	"github.com/galaktor/gorf24/reg/addr"
)

type RetransDelay byte

const (
	RETRANS_u250 = RetransDelay(iota)  // 0x0
	RETRANS_u500  // 0x1
	RETRANS_u750  // 0x2
	RETRANS_u1000 // 0x3
	RETRANS_u1250 // 0x4
	RETRANS_u1500 // 0x5
	RETRANS_u1750 // 0x6
	RETRANS_u2000 // 0x7
	RETRANS_u2250 // 0x8
	RETRANS_u2500 // 0x9
	RETRANS_u2750 // 0xA
	RETRANS_u3000 // 0xB
	RETRANS_u3250 // 0xC
	RETRANS_u3500 // 0xD
	RETRANS_u3750 // 0xE
	RETRANS_u4000 // 0xF
	
)

type SetupRetrans struct {
	R
}

func NewSetupRetrans(flags byte) *SetupRetrans {
	return &SetupRetrans{R{a: addr.SETUP_RETR, flags: flags}}
}

/* ARC
   set to 0 to disable */
func (r *SetupRetrans) SetCount(c uint8) error {
	if c > 15 {
		return errors.New(fmt.Sprintf("value out of legal range: %v. Only values 0 - 15 allowed.", c))
	}
	
	r.flags = (r.flags & 0xF0) | c
	return nil
}
func (r *SetupRetrans) GetCount() uint8 {
	return r.flags & 0x0F
}

/* ARD */
func (r *SetupRetrans) GetDelay() RetransDelay {
	return RetransDelay(r.flags >> 4)
}