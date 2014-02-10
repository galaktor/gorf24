package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type TransObserve struct {
	R
}

func NewTransObserve(flags byte) *TransObserve {
	return &TransObserve{R{a: addr.OBSERVE_TX, flags: flags}}
}