package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg/addr"
)

type RetransDelay byte

const (
	RETRANS_u250  = RetransDelay(iota) // 0x0
	RETRANS_u500                       // 0x1
	RETRANS_u750                       // 0x2
	RETRANS_u1000                      // 0x3
	RETRANS_u1250                      // 0x4
	RETRANS_u1500                      // 0x5
	RETRANS_u1750                      // 0x6
	RETRANS_u2000                      // 0x7
	RETRANS_u2250                      // 0x8
	RETRANS_u2500                      // 0x9
	RETRANS_u2750                      // 0xA
	RETRANS_u3000                      // 0xB
	RETRANS_u3250                      // 0xC
	RETRANS_u3500                      // 0xD
	RETRANS_u3750                      // 0xE
	RETRANS_u4000                      // 0xF

)

/* SETUP_RETR
Setup of Automatic Retransmission */
type SetupRetrans struct {
	R
}

func NewSetupRetrans(flags byte) *SetupRetrans {
	return &SetupRetrans{R{a: addr.SETUP_RETR, flags: flags}}
}

/* ARC (bits 3:0)
   Auto Retransmit Count
   ‘xxxx0000’ – Re-Transmit disabled
   ‘xxxx0001’ – Up to 1 Re-Transmit on fail of AA
   ......
   ‘xxxx1111’ – Up to 15 Re-Transmit on fail of AA */
func (r *SetupRetrans) SetCount(c uint8) error {
	if c > 15 {
		return errors.New(fmt.Sprintf("value out of legal range: %v. Only values 0 - 15 allowed.", c))
	}

	r.flags = (r.flags & 0xF0) | c
	return nil
}
func (r *SetupRetrans) GetCount() uint8 {
	return r.flags & 0x0F
}

/* ARD
   Auto Retransmit Delay
   ‘0000xxxx’ – Wait 250μS
   ‘0001xxxx’ – Wait 500μS
   ‘0010xxxx’ – Wait 750μS
   ........
   ‘1111xxxx’ – Wait 4000μS

   (Delay defined from end of transmission to start of
   next transmission)

   This is the time the PTX is waiting for an ACK packet
   before a retransmit is made. The PTX is in RX mode for
   250μS (500μS in 250kbps mode) to wait for address match.
   If the address match is detected, it stays in RX mode
   to the end of the packet, unless ARD elapses. Then it
   goes to standby-II mode for the rest of the specified
   ARD. After the ARD it goes to TX mode and then retrans-
   mits the packet.

   see constant values for the above times

   Please take care when setting this parameter. If the
   ACK payload is more than 15 byte in 2Mbps mode the
   ARD must be 500μS or more, if the ACK payload is more
   than 5byte in 1Mbps mode the ARD must be 500μS or more.
   In 250kbps mode (even when the payload is not in ACK)
   the ARD must be 500μS or more.*/
func (r *SetupRetrans) GetDelay() RetransDelay {
	return RetransDelay(r.flags >> 4)
}
func (r *SetupRetrans) SetDelay(d RetransDelay) {
	r.flags = (r.flags & 0x0F) | (byte(d) << 4)
}
