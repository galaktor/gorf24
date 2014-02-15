package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type FullXAddress struct {
	R

	flags XAddress
}

func NewFullXAddress(p pipe.P, flags XAddress) *FullXAddress {
	return &FullXAddress{R{addr.RX_ADDR(p),0},flags}
}

func (r *FullXAddress) Get() XAddress {
	return r.flags
}

func (r *FullXAddress) Set(a XAddress) {
	r.flags = a
}