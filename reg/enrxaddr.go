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
	switch p {
	case pipe.P0: e.flags = e.flags | 1
	case pipe.P1: e.flags = e.flags | 2
	case pipe.P2: e.flags = e.flags | 4
	case pipe.P3: e.flags = e.flags | 8
	case pipe.P4: e.flags = e.flags | 16
	case pipe.P5: e.flags = e.flags | 32
	}
}
func (e *EnabledRxAddresses) Disable(p pipe.P) {
	switch p {
	case pipe.P0: e.flags = e.flags & 0xFE
	case pipe.P1: e.flags = e.flags & 0xFD
	case pipe.P2: e.flags = e.flags & 0xFB
	case pipe.P3: e.flags = e.flags & 0xF7
	case pipe.P4: e.flags = e.flags & 0xEF
	case pipe.P5: e.flags = e.flags & 0xDF
	}
}