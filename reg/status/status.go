package status

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* STATUS
   Status Register (In parallel to the SPI command
   word applied on the MOSI pin, the STATUS register
   is shifted serially out on the MISO pin)
   bit 7 reserved */
type Status struct {
	reg.R
}

const res_mask = 0x7F

func NewStatus(flags byte) *Status {
	masked := flags & res_mask // reset reserved bits
	return &Status{reg.New(addr.STATUS, masked)}
}

/* TX_FULL (bit 0)
   TX FIFO full flag.
   1: TX FIFO full.
   0: Available locations in TX FIFO.

   tempted to return FifoUsage here like for FIFO_STATUS
   but the data here really *is* boolean */
func (s *Status) TxFull() bool {
	return (s.Byte() & 1) == 1
}

/* RX_P_NO (bits 3:1)
   Data pipe number for the payload available for
   reading from RX_FIFO
   000-101: Data Pipe Number
   110: Not Used
   111: RX FIFO Empty */
func (s *Status) RxPipeNumber() (pipe.P, error) {
	val := (s.Byte() >> 1) & 7

	switch {
	case val < 6:
		return pipe.P(val), nil
	case val == 7:
		return pipe.P(val), errors.New("Rx FIFO is empty.")
	}

	return pipe.P(val), errors.New(fmt.Sprintf("Invalid pipe number: %v", val))
}

/* when RX_P_NO bits are '111'
   tempted to return FifoUsage here like for FIFO_STATUS
   but the data here really *is* boolean */
func (s *Status) RxFifoEmpty() bool {
	p, _ := s.RxPipeNumber()
	return p == 7
}

/* MAX_RT (bit 4)
   Active low
   Maximum number of TX retransmits interrupt
   Write 1 to clear bit. If MAX_RT is asserted it must
   be cleared to enable further communication. */
func (s *Status) MaxTxRetransmits() bool {
	return (s.Byte() & 8) == 8
}
func (s *Status) ClearMaxTxRetransmits() {
	s.R.Set(s.Byte() | 8)
}

/* TX_DS (bit 5)
   Active low
   Data Sent TX FIFO interrupt. Asserted when
   packet transmitted on TX. If AUTO_ACK is acti-
   vated, this bit is set high only when ACK is
   received. */
func (s *Status) TxDataSent() bool {
	return (s.Byte() & 16) == 16
}
func (s *Status) ClearTxDataSent() {
	s.R.Set(s.Byte() | 16)
}

/* RX_DR (bit 6)
   Active low
   Data Ready RX FIFO interrupt. Asserted when
   new data arrives RX FIFO. */
func (s *Status) RxDataReady() bool {
	return (s.Byte() & 32) == 32
}
func (s *Status) ClearRxDataReady() {
	s.R.Set(s.Byte() | 32)
}
