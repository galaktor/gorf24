/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package config

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

type CrcLength byte

const (
	CRC_DISABLED CrcLength = iota
	CRC_8BIT
	CRC_16BIT
)

/* CONFIG
   Configuration Register
   bit 7 reserved */
type C struct {
	reg.R
}

func New() *C {
	return &C{reg.New(addr.CONFIG, 0x7F)} // 01111111
}

func NewWith(flags byte) *C {
	c := New()
	c.Set(flags)
	return c
}

/* PRIM_RX (bit 0)
   RX/TX control
   xxxxxxx0: PTX
   xxxxxxx1: PRX */
func (c *C) SetPrimaryReceiver() {
	c.R.Set(c.Get() | 1)
}
func (c *C) IsPrimaryReceiver() bool {
	return c.Get()&1 == 1
}
func (c *C) SetPrimaryTransmitter() {
	c.R.Set(c.Get() & 0xFE)
}
func (c *C) IsPrimaryTransmitter() bool {
	return c.Get()&1 == 0
}

/* PWR_UP (bit 1)
   xxxxxx0x: POWER DOWN
   xxxxxx1x: POWER UP */
func (c *C) SetPowerUp(up bool) {
	if up {
		c.R.Set(c.Get() | 2)
	} else {
		c.R.Set(c.Get() & 0xFD)
	}

}
func (c *C) IsPowerUp() bool {
	return c.Get()&2 == 2
}

/* CRCO (bit 2)
   CRC encoding scheme
   xxxxx0xx - 1 byte
   xxxxx1xx - 2 bytes */
func (c *C) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:
		c.Set(c.Get() & 0xFB)
	case CRC_16BIT:
		c.Set(c.Get() | 4)
	}
}
func (c *C) GetCrcLength() CrcLength {
	if c.Get()&4 == 4 {
		return CRC_16BIT
	} else {
		return CRC_8BIT
	}
}

/* EN_CRC (bit 3)
   Enable CRC. Forced high if one of the bits in the
   EN_AA is high

   xxxx0xxx - disabled
   xxxx1xxx - enabled */
func (c *C) SetCrcEnabled(enable bool) {
	if enable {
		c.R.Set(c.Get() | 8)
	} else {
		c.R.Set(c.Get() & 0xF7)
	}
}
func (c *C) IsCrcEnabled() bool {
	return c.Get()&8 == 8
}

/* MASK_MAX_RT (bit 4)
   Mask interrupt caused by MAX_RT

   xxx0xxxx: Reflect MAX_RT as active low interrupt on the IRQ pin
   xxx1xxxx: Interrupt not reflected on the IRQ pin  */
func (c *C) SetMaxRtInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Get() & 0xEF)
	} else {
		c.R.Set(c.Get() | 16)
	}
}
func (c *C) IsMaxRtInterruptEnabled() bool {
	return c.Get()&16 == 0
}

/* MASK_TX_DS (bit 5)
   Mask interrupt caused by TX_DS

   xx0xxxxx: Reflect TX_DS as active low interrupt on the IRQ pin
   xx1xxxxx: Interrupt not reflected on the IRQ pin*/
func (c *C) SetTxDsInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Get() & 0xDF)
	} else {
		c.R.Set(c.Get() | 32)
	}
}
func (c *C) IsTxDsInterruptEnabled() bool {
	return c.Get()&32 == 0
}

/* MASK_RX_DR (bit 6)
   Mask interrupt caused by RX_DR

   x0xxxxxx: Reflect RX_DR as active low interrupt on the IRQ pin
   x1xxxxxx: Interrupt not reflected on the IRQ pin */
func (c *C) SetRxDrInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Get() & 0xBF)
	} else {
		c.R.Set(c.Get() | 64)
	}
}
func (c *C) IsRxDrInterruptEnabled() bool {
	return c.Get()&64 == 0
}
