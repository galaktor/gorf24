package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

// only for internal use when creating actual RxAddr or TxAddr

type FullXAddress struct {
	R

	flags XAddress
}

func newFullXAddress(a addr.A, flags XAddress) *FullXAddress {
	return &FullXAddress{R{a,0},flags}
}

func (r *FullXAddress) Get() XAddress {
	return r.flags
}

func (r *FullXAddress) Set(a XAddress) {
	r.flags = a
}










