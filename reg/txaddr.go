package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* use to create XAddress for TX_ADDR

   TX_ADDR
   Transmit address. Used for a PTX device only.
   (LSByte is written first)
   Set RX_ADDR_P0 equal to this address to handle
   automatic acknowledge if this is a PTX device with
   Enhanced ShockBurstTM enabled. */
func NewTxAddress(flags XAddress) *FullXAddress {
	return newFullXAddress(addr.TX_ADDR, flags)
}
