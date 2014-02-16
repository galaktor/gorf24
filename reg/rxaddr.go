package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

func NewFullRxAddress(p pipe.P, flags XAddress) *FullXAddress {
	a := newFullXAddress(addr.RX_ADDR(p), flags)
	return a
}

func NewPartialRxAddress(p pipe.P, root *FullXAddress, lsb byte) *PartialXAddress {
	return newPartialXAddress(addr.RX_ADDR(p), root, lsb)
}















