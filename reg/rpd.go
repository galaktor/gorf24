package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

// bits 7:1 reserved

type ReceivedPowerDetector struct {
	R
}

func NewRPD(flags byte) *ReceivedPowerDetector {
	return &ReceivedPowerDetector{R{a: addr.RPD, flags: flags}}
}

/* RPD (bit 0) 
   triggers at received power levels above -64 dBm that
   are present in the RF channel you receive on. If the
   received power is less than -64 dBm, [sic] RDP = 0.*/
func (r *ReceivedPowerDetector) Triggered() bool {
	return r.flags & 1 == 1
}
