package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type FullPipeAddress struct {
	R

	flags PipeAddress
}

func NewFullPipeAddress(p pipe.P, flags PipeAddress) *FullPipeAddress {
	return &FullPipeAddress{R{addr.RX_ADDR(p),0},flags}
}

func (r *FullPipeAddress) Get() PipeAddress {
	return r.flags
}

func (r *FullPipeAddress) Set(a PipeAddress) {
	r.flags = a
}