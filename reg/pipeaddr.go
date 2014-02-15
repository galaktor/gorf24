package reg

import (

)

// TODO: move into pipe namespace?

type PipeAddress uint64

func NewPipeAddress(flags uint64) PipeAddress {
	// only use 5 least significant bytes
	// rest will be truncated
	return PipeAddress(flags & 0x000000FFFFFFFFFF)
}

func (a PipeAddress) Byte() uint64 {
	return uint64(a)
}
