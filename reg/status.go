package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

type Status struct {
	R
}

func NewStatus(flags byte) *Status {
	return &Status{R{a: addr.STATUS, flags: flags}}
}

/* TX_FULL (bit 0)
  TX FIFO full flag.
  1: TX FIFO full.
  0: Available locations in TX FIFO. */
func (s *Status) TxFull() bool {
	return (s.flags & 1) == 1
}

/* RX_P_NO (bits 3:1)
  Data pipe number for the payload available for
  reading from RX_FIFO 
  000-101: Data Pipe Number
  110: Not Used
  111: RX FIFO Empty */
func (s *Status) RxPipeNumber() (pipe.P,error) {
	val := (s.flags >> 1) & 7

	switch {
	case val < 6: return pipe.P(val),nil
	case val == 7: return pipe.P(val),errors.New("Rx FIFO is empty.")
	}

	return pipe.P(val),errors.New(fmt.Sprintf("Invalid pipe number: %v", val))
}

/* RX_P_NO bits from '000' up to '101' */
// now implied as error in RxPipeNumber()
/*func (s *Status) RxPipeNumberUsed() bool {
	return s.RxPipeNumber() < 6
}*/

/* RX_P_NO bits are '111' */
func (s *Status) RxFifoEmpty() bool {
	p,_ := s.RxPipeNumber()
	return p == 7
}

/* MAX_RT (bit 4)
  Maximum number of TX retransmits interrupt
  Write 1 to clear bit. If MAX_RT is asserted it must
  be cleared to enable further communication. */
func (s *Status) MaxTxRetransmits() bool {
	return (s.flags & 8) == 8
}

/* TX_DS (bit 5)
  Data Sent TX FIFO interrupt. Asserted when
  packet transmitted on TX. If AUTO_ACK is acti-
  vated, this bit is set high only when ACK is
  received. */
func (s *Status) TxDataSent() bool {
	return (s.flags & 16) == 16
}

/* RX_DR (bit 6)
  Data Ready RX FIFO interrupt. Asserted when
  new data arrives RX FIFO. */
func (s *Status) RxDataReady() bool {
	return (s.flags & 32) == 32
}