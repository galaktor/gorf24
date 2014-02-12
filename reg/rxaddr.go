package reg

import (

)

type RxAddress uint64

func NewRxAddress(flags uint64) RxAddress {
	// only use 5 least significant bytes
	// rest will be truncated
	return RxAddress(flags & 0x000000FFFFFFFFFF)
}

func (a RxAddress) Byte() uint64 {
	return uint64(a)
}
