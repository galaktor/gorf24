package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type TxAddress struct {
	R
}

// TODO: consider organizing differently with RxAddresses; share code or merge?

func NewTxAddress(flags byte) *TxAddress {
	return &TxAddress{R{a: addr.TX_ADDR, flags: flags}}
}