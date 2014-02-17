package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

// bit 7 reserved
// bits 3:2 reserved

type FifoUsage byte
const (
	FIFO_PARTIAL FifoUsage = iota // 0x0
	FIFO_EMPTY // 0x1
	FIFO_FULL  // 0x2
	FIFO_INVALID // 0x3
)

type FifoStatus struct {
	R
}

func NewFifoStatus(flags byte) *FifoStatus {
	return &FifoStatus{R{a: addr.FIFO_STATUS, flags: flags}}
}

/* RX_EMPTY (bit 0)
   RX_FULL  (bit 1) */
func (f *FifoStatus) Rx() FifoUsage {
	// xxxxxx00 -> partial
	// xxxxxx01 -> empty
	// xxxxxx10 -> full
	// xxxxxx11 -> empty AND full? INVALID!
	return FifoUsage(f.flags & 3)
}

/* TX_EMPTY (bit 4)
   TX_FULL  (bit 5) */
func (f *FifoStatus) Tx() FifoUsage {
	// xx00xxxx -> partial
	// xx01xxxx -> empty
	// xx10xxxx -> full
	// xx11xxxx -> empty AND full? INVALID!
	return FifoUsage((f.flags >> 4) & 3)
}
