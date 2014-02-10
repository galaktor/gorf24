package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type AutoAck struct {
	R
}

func NewAutoAck(flags byte) *AutoAck {
	return &AutoAck{R{a: addr.EN_AA, flags: flags}}
}