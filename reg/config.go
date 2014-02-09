package reg

import (

)

type ConfigReg struct {
	Register
}

func NewConfigReg(flags byte) *ConfigReg {
	return &ConfigReg{Register{a: Address(0), flags: flags}}
}

/* PRIM_RX */
func (c *ConfigReg) SetPrimaryReceiver() {
	c.flags = c.flags | 1
}
func (c *ConfigReg) IsPrimaryReceiver() bool {
	return c.flags&1 == 1
}
func (c *ConfigReg) SetPrimaryTransmitter() {
	c.flags = c.flags & 0xFE
}
func (c *ConfigReg) IsPrimaryTransmitter() bool {
	return c.flags&1 == 0
}

/* PWR_UP */
func (c *ConfigReg) SetPowerUp(up bool) {
	if up {
		c.flags = c.flags | 2
	} else {
		c.flags = c.flags & 0xFD
	}

}
func (c *ConfigReg) IsPowerUp() bool {
	return c.flags&2 == 2
}

/* CRCO */
func (c *ConfigReg) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:
		c.flags = c.flags & 0xFB
	case CRC_16BIT:
		c.flags = c.flags | 4
	}
}
func (c *ConfigReg) GetCrcLength() CrcLength {
	if c.flags&4 == 4 {
		return CRC_16BIT
	} else {
		return CRC_8BIT
	}
}

/* EN_CRC */
func (c *ConfigReg) SetCrcEnabled(enable bool) {
	if enable {
		c.flags = c.flags | 8
	} else {
		c.flags = c.flags & 0xF7
	}
}

/* MASK_MAX_RT */
func (c *ConfigReg) SetMaxRtInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xEF
	} else {
		c.flags = c.flags | 16
	}
}

/* MASK_TX_DS */
func (c *ConfigReg) SetTxDsInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xDF
	} else {
		c.flags = c.flags | 32
	}
}

/* MASK_RX_DR */
func (c *ConfigReg) SetRxDrInterruptEnabled(enable bool) {
	if enable {
		c.flags = c.flags & 0xBF
	} else {
		c.flags = c.flags | 64
	}
}