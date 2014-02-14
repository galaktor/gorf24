package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg/addr"
)

/* can only use two LSBits; 7:2 must be 000000 */

type AddrWidth byte

const (
	AW_3BYTES AddrWidth = 0x1
	AW_4BYTES AddrWidth = 0x2
	AW_5BYTES AddrWidth = 0x3
)

type AddrWidths struct {
	R
}

func NewAddrWidths(flags byte) *AddrWidths {
	return &AddrWidths{R{a: addr.SETUP_AW, flags: flags}}
}


/* AW
   value 00 is illegal; prevent it; return error maybe? */
func (a *AddrWidths) Set(w AddrWidth) error {
	if w == 0 || w & 0xFC > 0 {
		return errors.New(fmt.Sprintf("value outside of legal range: %v. only values 1 - 3 allowed.", w))
	}
	
	a.flags = byte(w)
	return nil
}
func (a *AddrWidths) Get() AddrWidth {
	return AddrWidth(a.flags)
}











