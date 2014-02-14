package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type PowerLevel byte

const (
	PA_MIN PowerLevel = iota // 0x0
	PA_LOW    // 0x1
	PA_MEDIUM // 0x2
	PA_MAX    // 0x3
)

// bit 0 obsolete/ignored

type RfSetup struct {
	R
}

func NewRfSetup(flags byte) *RfSetup {
	return &RfSetup{R{a: addr.RF_SETUP, flags: flags}}
}

/* RF_PWR */
func (s *RfSetup) Get() PowerLevel {
	return PowerLevel((s.flags & 0x6)>>1)
}








