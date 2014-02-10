package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type SetupRetrans struct {
	R
}

func NewSetupRetrans(flags byte) *SetupRetrans {
	return &SetupRetrans{R{a: addr.SETUP_RETR, flags: flags}}
}