package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type DynamicPayload struct {
	R
}

func NewDynamicPayload(flags byte) *DynamicPayload {
	return &DynamicPayload{R{a: addr.DYNPD, flags: flags}}
}
