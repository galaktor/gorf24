package reg

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* RPD
   Received Power Detector. This register is called
   CD (Carrier Detect) in the nRF24L01. The name is
   different in nRF24L01+ due to the different input
   power level threshold for this bit.
   bits 7:1 reserved */
type ReceivedPowerDetector struct {
	reg.R
}

func NewRPD(flags byte) *ReceivedPowerDetector {
	return &ReceivedPowerDetector{reg.New(addr.RPD, flags)}
}

/* RPD (bit 0)
   triggers at received power levels above -64 dBm that
   are present in the RF channel you receive on. If the
   received power is less than -64 dBm, [sic] RDP = 0.*/
func (r *ReceivedPowerDetector) Triggered() bool {
	return r.Byte()&1 == 1
}
