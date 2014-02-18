package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type FifoUsage byte

const (
	FIFO_PARTIAL FifoUsage = iota // 0x0
	FIFO_EMPTY                    // 0x1
	FIFO_FULL                     // 0x2
	FIFO_INVALID                  // 0x3
)

/* FIFO_STATUS
   FIFO Status Register
   bit 7 reserved
   bits 3:2 reserved */
type FifoStatus struct {
	R
}

func NewFifoStatus(flags byte) *FifoStatus {
	return &FifoStatus{R{a: addr.FIFO_STATUS, flags: flags}}
}

/* RX_EMPTY (bit 0)
   RX FIFO empty flag.
   xxxxxxx1: RX FIFO empty.
   xxxxxxx0: Data in RX FIFO.

   RX_FULL  (bit 1)
   RX FIFO full flag.
   xxxxxx1x: RX FIFO full.
   xxxxxx0x: Available locations in RX FIFO.

   both are used in combination here, as so:
   xxxxxx00 -> partial
   xxxxxx01 -> empty
   xxxxxx10 -> full
   xxxxxx11 -> empty AND full? INVALID!
*/
func (f *FifoStatus) Rx() FifoUsage {

	return FifoUsage(f.flags & 3)
}

/* TX_EMPTY (bit 4)
   TX FIFO empty flag.
   xxx1xxxx: TX FIFO empty.
   xxx0xxxx: Data in TX FIFO.

   TX_FULL  (bit 5)
   TX FIFO full flag.
   xx1xxxxx: TX FIFO full.
   xx0xxxxx: Available locations in TX FIFO.

   both are used in combination here, as so:
   xx00xxxx -> partial
   xx01xxxx -> empty
   xx10xxxx -> full
   xx11xxxx -> empty AND full? INVALID! */
func (f *FifoStatus) Tx() FifoUsage {

	return FifoUsage((f.flags >> 4) & 3)
}

/* TX_REUSE (bit 6)
   Used for a PTX device
   Pulse the rfce high for at least 10μs to Reuse
   last transmitted payload. TX payload reuse is
   active until W_TX_PAYLOAD or FLUSH TX is executed.
   TX_REUSE is set by the SPI command REUSE_TX_PL,
   and is reset by the SPI commands  W_TX_PAYLOAD
   or FLUSH TX

   x0xxxxxx -> disabled
   x1xxxxxx -> enabled */
func (f *FifoStatus) IsTxPayloadReuseEnabled() bool {
	return f.flags&0x40 == 0x40
}
