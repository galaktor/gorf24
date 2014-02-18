package reg

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

/* use to create 5byte XAddress for RX_ADDR_Px
   for a given pipe

   by spec only to be used for P0 and P1

   5 Bytes maximum
   length. (LSByte is written first. Write the number of
   bytes defined by SETUP_AW)*/
func NewFullRxAddress(p pipe.P, flags XAddress) *FullXAddress {
	a := newFullXAddress(addr.RX_ADDR(p), flags)
	return a
}

/* use to create 1byte (partial) XAddress for RX_ADDR_Px
   for a give pipe

   by spec only to be used for P2 - P5, which use P1 as root:

   Only LSB. MSBytes
   are equal to RX_ADDR_P1[39:8] */
func NewPartialRxAddress(p pipe.P, root *FullXAddress, lsb byte) *PartialXAddress {
	return newPartialXAddress(addr.RX_ADDR(p), root, lsb)
}
