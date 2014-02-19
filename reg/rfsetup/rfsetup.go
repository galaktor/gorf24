package rfsetup

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

type PowerLevel byte

const (
	PA_MIN    PowerLevel = iota // 0x0
	PA_LOW                      // 0x1
	PA_MEDIUM                   // 0x2
	PA_MAX                      // 0x3
)

type Datarate byte

const (
	RATE_1MBPS   Datarate = iota // 0x0
	RATE_2MBPS                   // 0x1
	RATE_250KBPS                 // 0x2
)

/* RF_SETUP
   RF Setup Register
   bit 6 reserved
   bit 4 only used in test (but exposed here anyway)
   bit 0 obsolete */
type R struct {
	reg.R
}

func New(flags byte) *R {
	return &R{reg.New(addr.RF_SETUP, flags)}
}

/* RF_PWR (bits 2:1)
   Set RF output power in TX mode
   'xxxxx00x' – -18dBm
   'xxxxx01x' – -12dBm
   'xxxxx10x' – -6dBm
   'xxxxx11x' –  0dBm */
func (s *R) GetPowerLevel() PowerLevel {
	return PowerLevel((s.Byte() & 0x6) >> 1)
}
func (s *R) SetPowerLevel(p PowerLevel) error {
	if p > 3 {
		return errors.New(fmt.Sprintf("Value out of legal range: %v. Allowed value from 0 -3.", p))
	}

	s.R.Set(s.Byte()&0xF9 | (byte(p) << 1))
	return nil
}

/* RF_DR_HIGH (bit 3)
   Select between the high speed data rates. This bit
   is don’t care if RF_DR_LOW is set.
   Encoding:
   [RF_DR_LOW, RF_DR_HIGH]:
   ‘xx0x0xxx’ – 1Mbps
   ‘xx0x1xxx’ – 2Mbps
   ‘xx1x0xxx’ – 250kbps
   ‘xx1x1xxx’ – Reserved

   RF_DR_LOW  (bit 5)
   Set RF Data Rate to 250kbps. See RF_DR_HIGH
   for encoding. */
func (s *R) GetDatarate() Datarate {
	var result Datarate

	switch val := s.Byte() & 0x28; val {
	case 0x0: // xx0x0xxx
		result = RATE_1MBPS
	case 0x8: // xx0x1xxx
		result = RATE_2MBPS
	case 0x20: // xx1x0xxx
		fallthrough // RF_DR_LOW set, so don't care
	case 0x28: // xx1x1xxx
		result = RATE_250KBPS
	}

	return result
}
func (s *R) SetDataRate(d Datarate) (err error) {
	err = nil

	switch d {
	case RATE_1MBPS: // xx0x0xxx
		s.R.Set(s.Byte() & 0xD7)
	case RATE_2MBPS: // xx0x1xxx
		s.R.Set((s.Byte() & 0xDF) | 8)
	case RATE_250KBPS: // xx1x0xxx
		s.R.Set((s.Byte() & 0xF7) | 32)
	default:
		err = errors.New(fmt.Sprintf("unsupported Datarate: %v. allowed values 0 - 2", d))
	}

	return
}

/* PLL_LOCK (bit 4)
   Force PLL lock signal. Only used in test

   spec not explicit on this, so I'm assuming:
   xxx0xxxx - disabled
   xxx1xxxx - enabled */
func (s *R) IsPllLockEnabled() bool {
	return s.Byte()&16 == 16
}
func (s *R) SetPllLock(enabled bool) {
	if enabled {
		s.R.Set(s.Byte() | 0x10)
	} else {
		s.R.Set(s.Byte() & 0xEF)
	}
}

/* CONT_WAVE
   Enables continuous carrier transmit when high. */
func (s *R) IsContinuousCarrierTransmitEnabled() bool {
	return s.Byte()&0x80 == 0x80 // 1xxxxxxx
}
func (s *R) SetContinuousCarrierTransmit(enabled bool) {
	if enabled {
		s.R.Set(s.Byte() | 0x80) // 1xxxxxxx
	} else {
		s.R.Set(s.Byte() & 0x7F) // 0xxxxxxx
	}
}
