package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type FullRxAddress struct {
	R

	flags RxAddress
}

func NewFullRxAddress(p pipe.P, flags RxAddress) *FullRxAddress {
	return &FullRxAddress{R{addr.RX_ADDR(p),0},flags}
}

func (r *FullRxAddress) Get() RxAddress {
	return r.flags
}

func (r *FullRxAddress) Set(a RxAddress) {
	r.flags = a
}