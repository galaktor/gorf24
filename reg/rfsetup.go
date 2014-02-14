package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg/addr"
)

type PowerLevel byte

const (
	PA_MIN PowerLevel = iota // 0x0
	PA_LOW    // 0x1
	PA_MEDIUM // 0x2
	PA_MAX    // 0x3
)

type Datarate byte

const (
	RATE_1MBPS Datarate = iota // 0x0
	RATE_2MBPS   // 0x1
	RATE_250KBPS // 0x2
)

// bit 0 obsolete/ignored
// bit 4 only used in test
// bit 6 reserved

type RfSetup struct {
	R
}

func NewRfSetup(flags byte) *RfSetup {
	return &RfSetup{R{a: addr.RF_SETUP, flags: flags}}
}

/* RF_PWR */
func (s *RfSetup) GetPowerLevel() PowerLevel {
	return PowerLevel((s.flags & 0x6)>>1)
}
func (s *RfSetup) SetPowerLevel(p PowerLevel) error {
	if p > 3 {
		return errors.New(fmt.Sprintf("Value out of legal range: %v. Allowed value from 0 -3.", p))
	}
	
	s.flags = s.flags & 0xF9 | (byte(p) << 1)
	return nil
}

/* RF_DR_LOW  (bit 5)
   RF_DR_HIGH (bit 3) */
func (s *RfSetup) GetDatarate() Datarate {
	var result Datarate

	switch val := s.flags & 0x28; val {
	case 0x0:  // xx0x0xxx
		result = RATE_1MBPS
	case 0x8:  // xx0x1xxx
		result = RATE_2MBPS
	case 0x20: // xx1x0xxx
		fallthrough
	case 0x28: // xx1x1xxx
		result = RATE_250KBPS
	}

	return result
}

/* PLL_LOCK 
   only used in test */
func (s *RfSetup) IsPllLockEnabled() bool {
	return s.flags & 16 == 16
}
func (s *RfSetup) SetPllLock(enabled bool) {
	if enabled {
		s.flags = s.flags | 0x10
	} else {
		s.flags = s.flags & 0xEF
	}
}