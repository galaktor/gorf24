package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type EnabledRxAddresses struct {
	R
}

func NewEnabledRxAddresses(flags byte) *EnabledRxAddresses {
	return &EnabledRxAddresses{R{addr.EN_RXADDR,0}}
}

func (e *EnabledRxAddresses) IsEnabled(p pipe.P) bool {
	return true // TODO
}

func (e *EnabledRxAddresses) SetEnabled(p pipe.P, enabled bool) {
	// TODO
}