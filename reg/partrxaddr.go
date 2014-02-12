package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type PartialRxAddress struct {
	R

	root *FullRxAddress
}

func NewPartialRxAddress(p pipe.P, root *FullRxAddress, lsb byte) *PartialRxAddress {
	return &PartialRxAddress{R{addr.RX_ADDR(p),lsb},root}
}


func (r *PartialRxAddress) Get() RxAddress {
	// use New method to force truncate
	return NewRxAddress((r.root.Get().Byte() << 8) | uint64(r.R.Byte()))
}

func (r *PartialRxAddress) Set(lsb byte) {
	r.R.flags = lsb
}
