package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type Config struct {
	R
}

func NewConfig(flags byte) *Config {
	return &Config{R{a: addr.CONFIG, flags: flags}}
}

/* PRIM_RX */
func (c *Config) SetPrimaryReceiver() {
	c.flags = c.flags | 1
}
func (c *Config) IsPrimaryReceiver() bool {
	return c.flags&1 == 1
}
func (c *Config) SetPrimaryTransmitter() {
	c.flags = c.flags & 0xFE
}
func (c *Config) IsPrimaryTransmitter() bool {
	return c.flags&1 == 0
}

/* PWR_UP */
func (c *Config) SetPowerUp(up bool) {
	if up {
		c.flags = c.flags | 2
	} else {
		c.flags = c.flags & 0xFD
	}

}
func (c *Config) IsPowerUp() bool {
	return c.flags&2 == 2
}

/* CRCO */
func (c *Config) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:
		c.flags = c.flags & 0xFB
	case CRC_16BIT:
		c.flags = c.flags | 4
	}
}
func (c *Config) GetCrcLength() CrcLength {
	if c.flags&4 == 4 {
		return CRC_16BIT
	} else {
		return CRC_8BIT
	}
}

/* EN_CRC */
func (c *Config) SetCrcEnabled(enable bool) {
	if enable {
		c.flags = c.flags | 8
	} else {
		c.flags = c.flags & 0xF7
	}
}
func (c *Config) IsCrcEnabled() bool {
	return c.flags & 8 == 8
}

/* MASK_MAX_RT */
func (c *Config) SetMaxRtInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xEF
	} else {
		c.flags = c.flags | 16
	}
}
func (c *Config) IsMaxRtInterruptEnabled() bool {
	return c.flags & 16 == 0
}


/* MASK_TX_DS */
func (c *Config) SetTxDsInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xDF
	} else {
		c.flags = c.flags | 32
	}
}
func (c *Config) IsTxDsInterruptEnabled() bool {
	return c.flags & 32 == 0
}

/* MASK_RX_DR */
func (c *Config) SetRxDrInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xBF
	} else {
		c.flags = c.flags | 64
	}
}
func (c *Config) IsRxDrInterruptEnabled() bool {
	return c.flags & 64 == 0
}
