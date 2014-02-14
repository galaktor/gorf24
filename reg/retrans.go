package reg

import (
	"errors"
	"fmt"
	
	"github.com/galaktor/gorf24/reg/addr"
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