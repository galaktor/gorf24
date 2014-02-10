package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type CarrierDetect struct {
	R
}

func NewCarrierDetect(flags byte) *CarrierDetect {
	return &CarrierDetect{R{a: addr.CD, flags: flags}}
}