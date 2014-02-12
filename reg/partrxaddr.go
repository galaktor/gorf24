package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type PartialRxAddress struct {
	R

	msb []byte // 4 bytes
}

func NewPartialRxAddress(p pipe.P, root *FullRxAddress, lsb byte) *PartialRxAddress {
	return &PartialRxAddress{R{addr.RX_ADDR(p),lsb},root.flags[0:4]} // slice is *reference* to parent byte array
}


func (r *PartialRxAddress) Byte() RxAddress {
	cp := [5]byte{}     // create copy to not spill address byte pointers
	copy(cp[:4], r.msb) // first 4 MSBytes from parent
	cp[4] = r.R.flags   // last  1 LSByte from own flags
	return cp           // copies again; TODO: have 'out' array preallocated in Partial, return (=copy) that as return
}

func (r *PartialRxAddress) Set(lsb byte) {
	r.R.flags = lsb
}