package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type RfSetup struct {
	R
}

func NewRfSetup(flags byte) *RfSetup {
	return &RfSetup{R{a: addr.RF_SETUP, flags: flags}}
}