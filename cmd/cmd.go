package cmd

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
)

type C byte

func (c C) Byte() byte {
	return byte(c)
}

var (
	NOP = C(util.B("11111111"))
	// R_REGISTER - is a function, see below
	// W_REGSITER - is a function, see below
	R_RX_PAYLOAD = C(util.B("01100001"))
	W_TX_PAYLOAD = C(util.B("10100000"))
	FLUSH_TX     = C(util.B("11100001"))
	FLUSH_RX     = C(util.B("11100010"))
	REUSE_TX_PL  = C(util.B("11100011"))
	ACTIVATE     = C(util.B("01010000"))
	// W_ACK_PAYLOAD - is a function, see below
	R_RX_PL_WID         = C(util.B("01100000"))
)

func R_REGISTER(r reg.R) C {
	return C(0x1F & r.Address())
}

func W_REGISTER(r reg.R) C {
	return C(0x20 | (0x1F & r.Address()))
}

func W_ACK_PAYLOAD(p pipe.P) C {
	return C(0xA8 | p)
}
