/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package cmd

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/pipe"
)

type C byte

func (c C) Byte() byte {
	return byte(c)
}

const (
	NOP                = C(0xFF) // 1111 1111
	// R_REGISTER - is a function, see below
	// W_REGSITER - is a function, see below
	R_RX_PAYLOAD       = C(0x61) // 0110 0001
	W_TX_PAYLOAD       = C(0xA0) // 1010 0000
	FLUSH_TX           = C(0xE1) // 1110 0001
	FLUSH_RX           = C(0xE2) // 1110 0010
	REUSE_TX_PL        = C(0xE3) // 1110 0011
	ACTIVATE           = C(0x50) // 0101 0000
	// W_ACK_PAYLOAD - is a function, see below
	R_RX_PL_WID        = C(0x60) // 0110 0000
	W_TX_PAYLOAD_NOACK = C(0xB0) // 1011 0000
)

func R_REGISTER(r reg.Register) C {
	return C(0x1F & r.Address())
}

func W_REGISTER(r reg.Register) C {
	return C(0x20 | (0x1F & r.Address()))
}

func W_ACK_PAYLOAD(p pipe.P) C {
	return C(0xA8 | p)
}







