package config

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

type CrcLength byte

const (
	CRC_DISABLED = iota
	CRC_8BIT
	CRC_16BIT
)

/* CONFIG
   Configuration Register
   bit 7 reserved */
type Config struct {
	reg.R
}

func NewConfig(flags byte) *Config {
	return &Config{reg.New(addr.CONFIG, flags)}
}

/* PRIM_RX (bit 0)
   RX/TX control
   xxxxxxx0: PTX
   xxxxxxx1: PRX */
func (c *Config) SetPrimaryReceiver() {
	c.R.Set(c.Byte() | 1)
}
func (c *Config) IsPrimaryReceiver() bool {
	return c.Byte()&1 == 1
}
func (c *Config) SetPrimaryTransmitter() {
	c.R.Set(c.Byte() & 0xFE)
}
func (c *Config) IsPrimaryTransmitter() bool {
	return c.Byte()&1 == 0
}

/* PWR_UP (bit 1)
   xxxxxx0x: POWER DOWN
   xxxxxx1x: POWER UP */
func (c *Config) SetPowerUp(up bool) {
	if up {
		c.R.Set(c.Byte() | 2)
	} else {
		c.R.Set(c.Byte() & 0xFD)
	}

}
func (c *Config) IsPowerUp() bool {
	return c.Byte()&2 == 2
}

/* CRCO (bit 2)
   CRC encoding scheme
   xxxxx0xx - 1 byte
   xxxxx1xx - 2 bytes */
func (c *Config) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:
		c.Set(c.Byte() & 0xFB)
	case CRC_16BIT:
		c.Set(c.Byte() | 4)
	}
}
func (c *Config) GetCrcLength() CrcLength {
	if c.Byte()&4 == 4 {
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
func (c *Config) SetCrcEnabled(enable bool) {
	if enable {
		c.R.Set(c.Byte() | 8)
	} else {
		c.R.Set(c.Byte() & 0xF7)
	}
}
func (c *Config) IsCrcEnabled() bool {
	return c.Byte()&8 == 8
}

/* MASK_MAX_RT (bit 4)
   Mask interrupt caused by MAX_RT

   xxx0xxxx: Reflect MAX_RT as active low interrupt on the IRQ pin
   xxx1xxxx: Interrupt not reflected on the IRQ pin  */
func (c *Config) SetMaxRtInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Byte() & 0xEF)
	} else {
		c.R.Set(c.Byte() | 16)
	}
}
func (c *Config) IsMaxRtInterruptEnabled() bool {
	return c.Byte()&16 == 0
}

/* MASK_TX_DS (bit 5)
   Mask interrupt caused by TX_DS

   xx0xxxxx: Reflect TX_DS as active low interrupt on the IRQ pin
   xx1xxxxx: Interrupt not reflected on the IRQ pin*/
func (c *Config) SetTxDsInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Byte() & 0xDF)
	} else {
		c.R.Set(c.Byte() | 32)
	}
}
func (c *Config) IsTxDsInterruptEnabled() bool {
	return c.Byte()&32 == 0
}

/* MASK_RX_DR (bit 6)
   Mask interrupt caused by RX_DR

   x0xxxxxx: Reflect RX_DR as active low interrupt on the IRQ pin
   x1xxxxxx: Interrupt not reflected on the IRQ pin */
func (c *Config) SetRxDrInterruptEnabled(enable bool) {
	if enable {
		c.R.Set(c.Byte() & 0xBF)
	} else {
		c.R.Set(c.Byte() | 64)
	}
}
func (c *Config) IsRxDrInterruptEnabled() bool {
	return c.Byte()&64 == 0
}
