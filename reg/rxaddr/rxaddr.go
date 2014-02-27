/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package rxaddr

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/reg/xaddr"
)

/* use to create 5byte XAddress for RX_ADDR_Px
   for a given pipe

   by spec only to be used for P0 and P1

   5 Bytes maximum
   length. (LSByte is written first. Write the number of
   bytes defined by SETUP_AW)*/
func NewFull(p pipe.P, flags xaddr.A) *xaddr.Full {
	a := xaddr.NewFull(addr.RX_ADDR(p), flags)
	return a
}

/* use to create 1byte (partial) XAddress for RX_ADDR_Px
   for a give pipe

   by spec only to be used for P2 - P5, which use P1 as root:

   Only LSB. MSBytes
   are equal to RX_ADDR_P1[39:8] */
func NewPartial(p pipe.P, root *xaddr.Full, lsb byte) *xaddr.Partial {
	return xaddr.NewPartialFrom(addr.RX_ADDR(p), root, lsb)
}
