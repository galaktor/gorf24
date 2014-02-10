package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type RxAddresses struct {
	R
}

// TODO: merge EN_RXADDR and RX_ADDR_PX together into one struct
// toggle enable and definition of addresses in methods here

func NewRxAddresses(flags byte) *RxAddresses {
	return &RxAddresses{R{a: addr.EN_RXADDR, flags: flags}}
}