package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type PartialXAddress struct {
	R

	root *FullXAddress
}

func NewPartialXAddress(p pipe.P, root *FullXAddress, lsb byte) *PartialXAddress {
	return &PartialXAddress{R{addr.RX_ADDR(p),lsb},root}
}


func (r *PartialXAddress) Get() XAddress {
	// use New method to force truncate
	return NewXAddress((r.root.Get().Byte() << 8) | uint64(r.R.Byte()))
}

func (r *PartialXAddress) Set(lsb byte) {
	r.R.flags = lsb
}
