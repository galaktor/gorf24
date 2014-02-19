package reg

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* FEATURE
   Feature Register
   bits 7:3 reserved */
type Feature struct {
	reg.R
}

func NewFeature(flags byte) *Feature {
	return &Feature{reg.New(addr.FEATURE, flags)}
}

/* EN_DYN_ACK (bit 0)
   Enables the W_TX_PAYLOAD_NOACK command
   xxxxxxx0 -> disabled
   xxxxxxx1 -> enabled */
func (f *Feature) IsDynamicAckEnabled() bool {
	return f.Byte()&0x01 == 0x01
}
func (f *Feature) SetDynamicAck(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x01)
	} else {
		f.R.Set(f.Byte() & 0xFE)
	}
}

/* EN_ACK_PAY (bit 1)
   Enables Payload with ACK

   If ACK packet payload is activated, ACK packets have
   dynamic payload lengths and the Dynamic PayloadLength
   feature should be enabled for pipe 0 on the PTX and PRX.
   This is to ensure that they receive the ACK packets
   with payloads. If the ACK payload is more than 15 byte
   in 2Mbps mode the ARD must be 500μS or more, and if the
   ACK payload is more than 5 byte in 1Mbps mode the ARD
   must be 500μS or more. In 250kbps mode (even when the
   payload is not in ACK) the ARD must be 500μS or more.

   xxxxxx0x -> disabled
   xxxxxx1x -> enabled */
func (f *Feature) IsPayloadWithAckEnabled() bool {
	return f.Byte()&0x02 == 0x02
}
func (f *Feature) SetPayloadWithAck(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x02)
	} else {
		f.R.Set(f.Byte() & 0xFD)
	}
}

/* EN_DPL (bit 2)
   Enables Dynamic Payload Length

   xxxxx0xx -> disabled
   xxxxx1xx -> enabled */
func (f *Feature) IsDynamicPayloadLengthEnabled() bool {
	return f.Byte()&0x04 == 0x04
}
func (f *Feature) SetDynamicPayloadLength(enabled bool) {
	if enabled {
		f.R.Set(f.Byte() | 0x04)
	} else {
		f.R.Set(f.Byte() & 0x7B)
	}
}
