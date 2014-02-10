package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type RfChannel struct {
	R
}

func NewRfChannel(flags byte) *RfChannel {
	return &RfChannel{R{a: addr.RF_CH, flags: flags}}
}