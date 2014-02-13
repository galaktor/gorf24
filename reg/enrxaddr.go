package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

type EnabledRxAddresses struct {
	R
}

func NewEnabledRxAddresses(flags byte) *EnabledRxAddresses {
	return &EnabledRxAddresses{R{addr.EN_RXADDR,flags}}
}

/* ERX_Px */
func (e *EnabledRxAddresses) IsEnabled(p pipe.P) bool {
	var result bool

	switch p {
	case pipe.P0: result = e.flags & 1 == 1
	case pipe.P1: result = e.flags & 2 == 2
	case pipe.P2: result = e.flags & 4 == 4
	case pipe.P3: result = e.flags & 8 == 8
	case pipe.P4: result = e.flags & 16 == 16
	case pipe.P5: result = e.flags & 32 == 32
	}

	return result
}
func (e *EnabledRxAddresses) Enable(p pipe.P) {
	// TODO
}
func (e *EnabledRxAddresses) Disable(p pipe.P) {
	// TODO
}