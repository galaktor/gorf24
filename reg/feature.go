package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

// bits 7:3 reserved

type Feature struct {
	R
}

func NewFeature(flags byte) *Feature {
	return &Feature{R{a: addr.FEATURE, flags: flags}}
}

/* EN_DYN_ACK (bit 0)
   Enables the W_TX_PAYLOAD_NOACK command */
func (f *Feature) IsDynamicAckEnabled() bool {
	return f.flags & 0x01 == 0x01
}
func (f *Feature) SetDynamicAck(enabled bool) {
	if enabled {
		f.flags |= 0x01
	} else {
		f.flags &= 0xFE
	}
}

/* EN_ACK_PAY (bit 1)
   Enables Payload with ACK */
func (f *Feature) IsPayloadWithAckEnabled() bool {
	return f.flags & 0x02 == 0x02
}
func (f *Feature) SetPayloadWithAck(enabled bool) {
	if enabled {
		f.flags |= 0x02
	} else {
		f.flags &= 0xFD
	}
}