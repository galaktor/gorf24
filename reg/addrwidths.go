package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type AddrWidths struct {
	R
}

func NewAddrWidths(flags byte) *AddrWidths {
	return &AddrWidths{R{a: addr.SETUP_AW, flags: flags}}
}