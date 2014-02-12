package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type ReceivedPowerDetector struct {
	R
}

func NewRPD(flags byte) *ReceivedPowerDetector {
	return &ReceivedPowerDetector{R{a: addr.RPD, flags: flags}}
}