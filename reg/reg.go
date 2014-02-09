package reg

import (
	"github.com/galaktor/gorf24/pipe"
)

type Register struct {
	a Address
	flags byte
}

func (c *Register) Byte() byte {
	return c.flags
}

var (
	/* Configuration Register */
	REG_CONFIG = Address(0)

	/* Enhanced ShockBurst
	   Enable ‘Auto Acknowledgment’ Function Dis-
	   able this functionality to be compatible with
	   nRF2401 */
	REG_EN_AA = Address(0x1)

	/* Enabled RX Addresses */
	REG_EN_RXADDR = Address(0x2)

	/* Setup of Address Widths
	(common for all data pipes) */
	REG_SETUP_AW = Address(0x3)

	/* Setup of Automatic Retransmission */
	REG_SETUP_RETR = Address(0x4)

	/* RF Channel */
	REG_RF_CH = Address(0x5)

	/* RF Setup Register */
	REG_RF_SETUP = Address(0x6)

	/* Status Register (In parallel to the SPI command
	           word applied on the MOSI pin, the STATUS reg-
		   ister is shifted serially out on the MISO pin)
		   (bit 7 reserved)*/
	REG_STATUS = Address(0x7)

	/* Transmit observe register */
	REG_OBSERVE_TX = Address(0x8)

	/* Carrier detect
	   (bits 7:1 reserved) */
	REG_CD = Address(0x9)

	// REG_RX_ADDR is a function - see below

	/* Transmit address. Used for a PTX device only.
	   (LSByte is written first)
	   Set RX_ADDR_P0 equal to this address to han-
	   dle automatic acknowledge if this is a PTX
	   device with Enhanced ShockBurstTM enabled. */
	REG_TX_ADDR = Address(0x10)

	// REG_RX_PW is a function - see below

	/* FIFO Status Register
	   (bits 7,3:2 reserved)*/
	REG_FIFO_STATUS = Address(0x17)

	/* n/a = separate SPI commands
	   REG_ACK_PLD
	   REG_TX_PLD
	   REG_RX_PLD
	*/

	/* Enable dynamic payload length
	   (bits 7:6 reserved) */
	REG_DYNPD = Address(0x1C)

	/* Feature register
	   To activate this feature use the ACTIVATE SPI command followed by
	   data 0x73. The corresponding bits in the FEATURE register must
	   be set.
	   (bits 7:3 reserved) */
	REG_FEATURE = Address(0x1D)
)

/* Receive address data pipe 0. 5 Bytes maximum
   length. (LSByte is written first. Write the number
   of bytes defined by SETUP_AW) */
func REG_RX_ADDR(p pipe.P) Address {
	return Address(0x0A + p)
} 

/* Number of bytes in RX payload in data pipe 0-5
   (1 to 32 bytes).
   0 Pipe not used
   1 = 1 byte
   ...
   32 = 32 bytes
   (bits 7:6 reserved)*/
func REG_RX_PW(p pipe.P) Address {
	return Address(0x11 + p)
}
