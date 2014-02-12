package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

// values up to 32; truncate above 6th bit?

type RxPayloadWidth struct {
	R
}

func NewRxPayloadWidth(p pipe.P, flags byte) *RxPayloadWidth {
	return &RxPayloadWidth{R{a: addr.RX_PW(p), flags: flags}}
}