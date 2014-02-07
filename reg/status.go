package reg

// TODO: fix to be status

type Status byte

/* TX_FULL (bit 0)
  TX FIFO full flag.
  1: TX FIFO full.
  0: Available locations in TX FIFO. */
func (s Status) TxFull() bool {
	return (s & 1) == 1
}

/* RX_P_NO (bits 3:1)
  Data pipe number for the payload available for
  reading from RX_FIFO 
  000-101: Data Pipe Number
  110: Not Used
  111: RX FIFO Empty */
func (s Status) RxPipeNumber() uint8 {
	return uint8((s >> 1) & 7)
}

/* RX_P_NO bits from '000' up to '101' */
func (s Status) RxPipeNumberUsed() bool {
	return s.RxPipeNumber() < 6
}

/* RX_P_NO bits are '111' */
func (s Status) RxFifoEmpty() bool {
	return s.RxPipeNumber() == 7
}

/* MAX_RT (bit 4)
  Maximum number of TX retransmits interrupt
  Write 1 to clear bit. If MAX_RT is asserted it must
  be cleared to enable further communication. */
func (s Status) MaxTxRetransmits() bool {
	return (s & 8) == 8
}

/* TX_DS (bit 5)
  Data Sent TX FIFO interrupt. Asserted when
  packet transmitted on TX. If AUTO_ACK is acti-
  vated, this bit is set high only when ACK is
  received. */
func (s Status) TxDataSent() bool {
	return (s & 16) == 16
}

/* RX_DR (bit 6)
  Data Ready RX FIFO interrupt. Asserted when
  new data arrives RX FIFO. */
func (s Status) RxDataReady() bool {
	return (s & 32) == 32
}