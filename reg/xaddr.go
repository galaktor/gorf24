package reg

import (

)

// used for Rx and Tx addresses; hence 'x' address
type XAddress uint64

func NewXAddress(flags uint64) XAddress {
	// only use 5 least significant bytes
	// rest will be truncated
	return XAddress(flags & 0x000000FFFFFFFFFF)
}

func (a XAddress) Byte() uint64 {
	return uint64(a)
}
