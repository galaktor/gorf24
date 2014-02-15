package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type PartialPipeAddress struct {
	R

	root *FullPipeAddress
}

func NewPartialPipeAddress(p pipe.P, root *FullPipeAddress, lsb byte) *PartialPipeAddress {
	return &PartialPipeAddress{R{addr.RX_ADDR(p),lsb},root}
}


func (r *PartialPipeAddress) Get() PipeAddress {
	// use New method to force truncate
	return NewPipeAddress((r.root.Get().Byte() << 8) | uint64(r.R.Byte()))
}

func (r *PartialPipeAddress) Set(lsb byte) {
	r.R.flags = lsb
}
