package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type FifoStatus struct {
	R
}

func NewFifoStatus(flags byte) *FifoStatus {
	return &FifoStatus{R{a: addr.FIFO_STATUS, flags: flags}}
}
