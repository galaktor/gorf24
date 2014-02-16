package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type PartialXAddress struct {
	R

	root *FullXAddress
}

func newPartialXAddress(a addr.A, root *FullXAddress, lsb byte) *PartialXAddress {
	return &PartialXAddress{R{a,lsb},root}
}

func (r *PartialXAddress) Get() XAddress {
	// use New method to force truncate
	return NewXAddress((r.root.Get().Byte() << 8) | uint64(r.R.Byte()))
}

func (r *PartialXAddress) Set(lsb byte) {
	r.R.flags = lsb
}
