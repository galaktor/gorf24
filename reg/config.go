package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* CONFIG
   Configuration Register
   bit 7 reserved */
type Config struct {
	R
}

func NewConfig(flags byte) *Config {
	return &Config{R{a: addr.CONFIG, flags: flags}}
}

/* PRIM_RX (bit 0)
   RX/TX control
   1: PRX, 0: PTX */
func (c *Config) SetPrimaryReceiver() {
	c.flags = c.flags | 1
}
func (c *Config) IsPrimaryReceiver() bool {
	return c.flags&1 == 1
}
func (c *Config) SetPrimaryTransmitter() {
	c.flags &= 0xFE
}
func (c *Config) IsPrimaryTransmitter() bool {
	return c.flags&1 == 0
}

/* PWR_UP (bit 1)
   1: POWER UP, 0:POWER DOWN */
func (c *Config) SetPowerUp(up bool) {
	if up {
		c.flags |= 2
	} else {
		c.flags &= 0xFD
	}

}
func (c *Config) IsPowerUp() bool {
	return c.flags&2 == 2
}

/* CRCO (bit 2)
   CRC encoding scheme
   '0' - 1 byte
   '1' – 2 bytes */
func (c *Config) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:
		c.flags &= 0xFB
	case CRC_16BIT:
		c.flags |= 4
	}
}
func (c *Config) GetCrcLength() CrcLength {
	if c.flags&4 == 4 {
		return CRC_16BIT
	} else {
		return CRC_8BIT
	}
}

/* EN_CRC (bit 3)
   Enable CRC. Forced high if one of the bits in the
   EN_AA is high */
func (c *Config) SetCrcEnabled(enable bool) {
	if enable {
		c.flags |= 8
	} else {
		c.flags &= 0xF7
	}
}
func (c *Config) IsCrcEnabled() bool {
	return c.flags&8 == 8
}

/* MASK_MAX_RT (bit 4)
   Mask interrupt caused by MAX_RT
   1: Interrupt not reflected on the IRQ pin
   0: Reflect MAX_RT as active low interrupt on the
   IRQ pin */
func (c *Config) SetMaxRtInterruptEnabled(enable bool) {
	if enable {
		c.flags &= 0xEF
	} else {
		c.flags |= 16
	}
}
func (c *Config) IsMaxRtInterruptEnabled() bool {
	return c.flags&16 == 0
}

/* MASK_TX_DS (bit 5)
   Mask interrupt caused by TX_DS
   1: Interrupt not reflected on the IRQ pin
   0: Reflect TX_DS as active low interrupt on the
   IRQ pin */
func (c *Config) SetTxDsInterruptEnabled(enable bool) {
	if enable {
		c.flags &= 0xDF
	} else {
		c.flags |= 32
	}
}
func (c *Config) IsTxDsInterruptEnabled() bool {
	return c.flags&32 == 0
}

/* MASK_RX_DR (bit 6)
   Mask interrupt caused by RX_DR
   1: Interrupt not reflected on the IRQ pin
   0: Reflect RX_DR as active low interrupt on the
   IRQ pin */
func (c *Config) SetRxDrInterruptEnabled(enable bool) {
	if enable {
		c.flags &= 0xBF
	} else {
		c.flags |= 64
	}
}
func (c *Config) IsRxDrInterruptEnabled() bool {
	return c.flags&64 == 0
}
