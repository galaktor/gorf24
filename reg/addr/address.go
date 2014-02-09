package addr

import (
	"github.com/galaktor/gorf24/pipe"
)

type A byte

func (a A) Byte() byte {
	return byte(a)
}

/*
Note: Addresses 18 to 1B are reserved for test purposes, altering them will make the chip malfunc-
tion.
*/

const (
	/* Configuration Register */
	CONFIG = A(0)

	/* Enhanced ShockBurst
	   Enable ‘Auto Acknowledgment’ Function Dis-
	   able this functionality to be compatible with
	   nRF2401 */
	EN_AA = A(0x1)

	/* Enabled RX Addresses */
	EN_RXADDR = A(0x2)

	/* Setup of Address Widths
	(common for all data pipes) */
	SETUP_AW = A(0x3)

	/* Setup of Automatic Retransmission */
	SETUP_RETR = A(0x4)

	/* RF Channel */
	RF_CH = A(0x5)

	/* RF Setup Register */
	RF_SETUP = A(0x6)

	/* Status Register (In parallel to the SPI command
	           word applied on the MOSI pin, the STATUS reg-
		   ister is shifted serially out on the MISO pin)
		   (bit 7 reserved)*/
	STATUS = A(0x7)

	/* Transmit observe register */
	OBSERVE_TX = A(0x8)

	/* Carrier detect
	   (bits 7:1 reserved) */
	CD = A(0x9)

	// RX_ADDR is a function - see below

	/* Transmit address. Used for a PTX device only.
	   (LSByte is written first)
	   Set RX_ADDR_P0 equal to this address to han-
	   dle automatic acknowledge if this is a PTX
	   device with Enhanced ShockBurstTM enabled. */
	TX_ADDR = A(0x10)

	// RX_PW is a function - see below

	/* FIFO Status Register
	   (bits 7,3:2 reserved)*/
	FIFO_STATUS = A(0x17)

	/* n/a = separate SPI commands
	   ACK_PLD
	   TX_PLD
	   RX_PLD
	*/

	/* Enable dynamic payload length
	   (bits 7:6 reserved) */
	DYNPD = A(0x1C)

	/* Feature register
	   To activate this feature use the ACTIVATE SPI command followed by
	   data 0x73. The corresponding bits in the FEATURE register must
	   be set.
	   (bits 7:3 reserved) */
	FEATURE = A(0x1D)
)

/* Receive address data pipe 0. 5 Bytes maximum
   length. (LSByte is written first. Write the number
   of bytes defined by SETUP_AW) */
func RX_ADDR(p pipe.P) A {
	return A(0x0A + p)
} 

/* Number of bytes in RX payload in data pipe 0-5
   (1 to 32 bytes).
   0 Pipe not used
   1 = 1 byte
   ...
   32 = 32 bytes
   (bits 7:6 reserved)*/
func RX_PW(p pipe.P) A {
	return A(0x11 + p)
}