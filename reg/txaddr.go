package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

func NewTxAddress(flags XAddress) *FullXAddress {
	return newFullXAddress(addr.TX_ADDR,flags)
}
















