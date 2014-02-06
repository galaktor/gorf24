/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The MIT License (see README and LICENSE files) */

package gorf24

import (
	"strconv"
	"time"

	"./gpio"
	"./spi"
)

type PA_DBM byte

const (
	PA_MIN PA_DBM = iota
	PA_LOW
	PA_HIGH
	PA_MAX
	PA_ERROR // what is this for?
)

type DATARATE byte

const (
	RATE_1MBPS DATARATE = iota
	RATE_2MBPS
	RATE_250KBPS
)

type CrcLength byte

const (
	CRC_DISABLED = iota
	CRC_8BIT
	CRC_16BIT
)

const RF24_PAYLOAD_SIZE = 32

type R struct {
	spi    *spi.SPI
	ce     *gpio.Pin
	csn    *gpio.Pin
	status Status
}

func New(spidevice string, spispeed uint32, cepin, csnpin uint8) (r *R, err error) {
	r = &R{}

	r.spi, err = spi.New(spidevice, 0, 8, spi.SPD_02MHz)
	if err != nil {
		return
	}

	r.ce, err = gpio.Open(cepin, gpio.OUT)
	if err != nil {
		return
	}

	r.csn, err = gpio.Open(csnpin, gpio.OUT)
	if err != nil {
		return
	}

	r.ce.SetLow()
	r.csn.SetHigh()

	// ** FROM RF24.cpp **
	// Must allow the radio time to settle else configuration bits will not necessarily stick.
	// This is actually only required following power up but some settling time also appears to
	// be required after resets too. For full coverage, we'll always assume the worst.
	// Enabling 16b CRC is by far the most obvious case if the wrong timing is used - or skipped.
	// Technically we require 4.5ms + 14us as a worst case. We'll just call it 5ms for good measure.
	// WARNING: Delay is based on P-variant whereby non-P *may* require different timing.
	<-time.After(5 * time.Millisecond)

	return
}

func (r *R) Status() Status {
	return r.status
}

// sends Command, then buf byte-by-byte over SPI
// if buf is null, sends only command
// WARNING: destructive - overwrites content of buf while pumping
func (r *R) spiPump(c Command, buf []byte) error {
	r.csn.SetLow()
	defer r.csn.SetHigh()

	// send cmd first
	s, err := r.spi.Transfer(c.Byte())
	if err != nil {
		return err
	}
	r.status = Status(s)

	if buf != nil {
		// pump buf data, overwriting content with returned date
		// RF24 SPI does LSByte first, so iterate backward
		for n := len(buf); n >= 0; n-- {
			buf[n], err = r.spi.Transfer(buf[n])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ?TODO: cap data sizes allowed for []byte?

// ?TODO: use buffer(s) pre-allocated on R?

/*
Read command and status registers. AAAAA =
5 bit Register Map Address
*/
func (r *R) readRegister(reg Register, buf []byte) error {
	return r.spiPump(CMD_R_REGISTER(reg), buf)
}

/*
Write command and status registers. AAAAA = 5
bit Register Map Address
Executable in power down or standby modes
only.
*/
func (r *R) writeRegister(reg Register, buf []byte) error {
	return r.spiPump(CMD_W_REGISTER(reg), buf)
}

/*
Read RX-payload: 1 – 32 bytes. A read operation
always starts at byte 0. Payload is deleted from
FIFO after it is read. Used in RX mode
*/
func (r *R) readPayload(buf []byte) error {
	// ?TODO: set to RX mode?
	return r.spiPump(CMD_R_RX_PAYLOAD, buf)
}

/*
Write TX-payload: 1 – 32 bytes. A write operation
always starts at byte 0 used in TX payload.
*/
func (r *R) writePayload(buf []byte) error {
	// ?TODO: set to TX mode?
	return r.spiPump(CMD_W_TX_PAYLOAD, buf)
}

// ?TODO: enum for modes, MD_RX and MD_TX?

/*
Flush TX FIFO, used in TX mode
*/
func (r *R) flushTx() error {
	// ?TODO: enforce/check mode?
	return r.spiPump(CMD_FLUSH_TX, nil)
}

/*
Flush RX FIFO, used in RX mode
Should not be executed during transmission of
acknowledge, that is, acknowledge package will
not be completed.
*/
func (r *R) flushRx() error {
	// ?TODO: enforce/check mode?
	// from spec:
	//   Should not be executed during transmission of
	//   acknowledge, that is, acknowledge package will
	//   not be completed.
	return r.spiPump(CMD_FLUSH_RX, nil)
}

/*
No Operation. Might be used to read the STATUS
register
*/
func (r *R) refreshStatus() (Status, error) {
	// spiPump will update status on every cmd sent
	err := r.spiPump(CMD_NOP, nil)
	return r.status, err
}

/*
This write command followed by data 0x73 acti-
vates the following features:
• R_RX_PL_WID
• W_ACK_PAYLOAD
• W_TX_PAYLOAD_NOACK
A new ACTIVATE command with the same data
deactivates them again. This is executable in
power down or stand by modes only.
The R_RX_PL_WID, W_ACK_PAYLOAD, and
W_TX_PAYLOAD_NOACK features registers are
initially in a deactivated state; a write has no
effect, a read only results in zeros on MISO. To
activate these registers, use the ACTIVATE com-
mand followed by data 0x73. Then they can be
accessed as any other register in nRF24L01. Use
the same command and data to deactivate the
registers again.
*/
func (r *R) toggleActivate() error {
	// TODO: keep activated bool state, make de/activate() funcs
	return r.spiPump(CMD_ACTIVATE, []byte{0x73})
}

/***** pipe.go *****/
// 0 through 5, giving 6 data pipe ids
type Pipe byte

/***** EOF pipe.go *****/

/***** reg.go *****/

/*
Note: Addresses 18 to 1B are reserved for test purposes, altering them will make the chip malfunc-
tion.
*/

type Register byte

func (r Register) Byte() byte {
	return byte(r)
}

var (
	/* Configuration Register */
	REG_CONFIG = Register(0)

	/* Enhanced ShockBurst
	   Enable ‘Auto Acknowledgment’ Function Dis-
	   able this functionality to be compatible with
	   nRF2401 */
	REG_EN_AA = Register(0x1)

	/* Enabled RX Addresses */
	REG_EN_RXADDR = Register(0x2)

	/* Setup of Address Widths
	(common for all data pipes) */
	REG_SETUP_AW = Register(0x3)

	/* Setup of Automatic Retransmission */
	REG_SETUP_RETR = Register(0x4)

	/* RF Channel */
	REG_RF_CH = Register(0x5)

	/* RF Setup Register */
	REG_RF_SETUP = Register(0x6)

	/* Status Register (In parallel to the SPI command
	           word applied on the MOSI pin, the STATUS reg-
		   ister is shifted serially out on the MISO pin)
		   (bit 7 reserved)*/
	REG_STATUS = Register(0x7)

	/* Transmit observe register */
	REG_OBSERVE_TX = Register(0x8)

	/* Carrier detect
	   (bits 7:1 reserved) */
	REG_CD = Register(0x9)

	/* Receive address data pipe 0. 5 Bytes maximum
	   length. (LSByte is written first. Write the number
	   of bytes defined by SETUP_AW) */
	REG_RX_ADDR_P0 = Register(0x0A)

	/* Receive address data pipe 1. 5 Bytes maximum
	   length. (LSByte is written first. Write the number
	   of bytes defined by SETUP_AW) */
	REG_RX_ADDR_P1 = Register(0x0B)

	/* Receive address data pipe 2. Only LSB. MSBy-
	   tes is equal to RX_ADDR_P1[39:8] */
	REG_RX_ADDR_P2 = Register(0x0C)

	/* Receive address data pipe 3. Only LSB. MSBy-
	   tes is equal to RX_ADDR_P1[39:8] */
	REG_RX_ADDR_P3 = Register(0x0D)

	/* Receive address data pipe 4. Only LSB. MSBy-
	   tes is equal to RX_ADDR_P1[39:8] */
	REG_RX_ADDR_P4 = Register(0x0E)

	/* Receive address data pipe 5. Only LSB. MSBy-
	   tes is equal to RX_ADDR_P1[39:8] */
	REG_RX_ADDR_P5 = Register(0x0F)

	/* Transmit address. Used for a PTX device only.
	   (LSByte is written first)
	   Set RX_ADDR_P0 equal to this address to han-
	   dle automatic acknowledge if this is a PTX
	   device with Enhanced ShockBurstTM enabled. */
	REG_TX_ADDR = Register(0x10)

	/* Number of bytes in RX payload in data pipe 0-5
	   (1 to 32 bytes).
	   0 Pipe not used
	   1 = 1 byte
	   ...
	   32 = 32 bytes
	   (bits 7:6 reserved)*/
	REG_RX_PW_P0 = Register(0x11)
	REG_RX_PW_P1 = Register(0x12)
	REG_RX_PW_P2 = Register(0x13)
	REG_RX_PW_P3 = Register(0x14)
	REG_RX_PW_P4 = Register(0x15)
	REG_RX_PW_P5 = Register(0x16)

	/* FIFO Status Register
	   (bits 7,3:2 reserved)*/
	REG_FIFO_STATUS = Register(0x17)

	/* n/a = separate SPI commands
	   REG_ACK_PLD
	   REG_TX_PLD
	   REG_RX_PLD
	*/

	/* Enable dynamic payload length
	   (bits 7:6 reserved) */
	REG_DYNPD = Register(0x1C)

	/* Feature register
	   To activate this feature use the ACTIVATE SPI command followed by
	   data 0x73. The corresponding bits in the FEATURE register must
	   be set.
	   (bits 7:3 reserved) */
	REG_FEATURE = Register(0x1D)
)

type ConfigReg struct {
	address Register
	flags   byte
}

func NewConfigReg(flags byte) *ConfigReg {
	return &ConfigReg{address: Register(0), flags: flags}
}

func (c *ConfigReg) Byte() byte {
	return c.flags
}

/* PRIM_RX */
func (c *ConfigReg) SetPrimaryReceiver() {
	c.flags =  c.flags | 1
}
func (c *ConfigReg) IsPrimaryReceiver() bool {
	return c.flags & 1 == 1
}
func (c *ConfigReg) SetPrimaryTransmitter() {
	c.flags = c.flags & 0xFE
}
func (c *ConfigReg) IsPrimaryTransmitter() bool {
	return c.flags & 1 == 0 
}

/* PWR_UP */
func (c *ConfigReg) SetPowerUp() {
	c.flags = c.flags | 2
}
func (c *ConfigReg) IsPowerUp() bool {
	return c.flags & 2 == 2
}
func (c *ConfigReg) SetPowerDown() {
	c.flags = c.flags & 0xFD
}
func (c *ConfigReg) IsPowerDown() bool {
	return c.flags & 2 == 0
}

/* CRCO */
func (c *ConfigReg) SetCrcLength(l CrcLength) {
	switch l {
	case CRC_8BIT:  c.flags = c.flags & 0xFB
	case CRC_16BIT: c.flags = c.flags | 4
	}
}
func (c *ConfigReg) GetCrcLength() CrcLength {
	if c.flags & 4 == 4 {
		return CRC_16BIT
	} else {
		return CRC_8BIT
	}
}

/***** EOF reg.go *****/

/***** cmd.go *****/
type Command byte

func (c Command) Byte() byte {
	return byte(c)
}

var (
	CMD_NOP = Command(B("11111111"))
	// CMD_R_REGISTER - is a function, see below
	// CMD_W_REGSITER - is a function, see below
	CMD_R_RX_PAYLOAD = Command(B("01100001"))
	CMD_W_TX_PAYLOAD = Command(B("10100000"))
	CMD_FLUSH_TX     = Command(B("11100001"))
	CMD_FLUSH_RX     = Command(B("11100010"))
	CMD_REUSE_TX_PL  = Command(B("11100011"))
	CMD_ACTIVATE     = Command(B("01010000"))
	// CMD_W_ACK_PAYLOAD - is a function, see below
	CMD_W_ACK_PAYLOAD_PIPE0 = Command(B("10101000"))
	CMD_W_ACK_PAYLOAD_PIPE1 = Command(B("10101001"))
	CMD_W_ACK_PAYLOAD_PIPE2 = Command(B("10101010"))
	CMD_W_ACK_PAYLOAD_PIPE3 = Command(B("10101011"))
	CMD_W_ACK_PAYLOAD_PIPE4 = Command(B("10101100"))
	CMD_W_ACK_PAYLOAD_PIPE5 = Command(B("10101101"))
	CMD_R_RX_PL_WID         = Command(B("01100000"))
)

func CMD_R_REGISTER(r Register) Command {
	return Command(0x1F & r)
}

func CMD_W_REGISTER(r Register) Command {
	return Command(0x20 | (0x1F & r))
}

func CMD_W_ACK_PAYLOAD(p Pipe) Command {
	return Command(0xA8 | p)
}

/***** EOF cmd.go *****/

/***** util.go *****/
/*
 parse a string representation of bits into
 a byte; for easier testing
*/
func B(bits string) byte {
	i, _ := strconv.ParseUint(bits, 2, 8)
	return byte(i)
}

/***** EOF util.go *****/

/**********************/
/***** OLD STUFF! *****/

/*
func (r *R) Delete() {
	C.rf24_delete(r.cptr)
	r.cptr = nil
}

func (r *R) Begin() {
	C.rf24_begin(r.cptr)
}

func (r *R) ResetCfg() {
	C.rf24_resetcfg(r.cptr)
}

func (r *R) StartListening() {
	C.rf24_startListening(r.cptr)
}

func (r *R) StopListening() {
	C.rf24_stopListening(r.cptr)
}

// TODO: implement Reader/Writer compatible interfaces
func (r *R) Write(data []byte, length uint8) bool {
	return gobool(C.rf24_write(r.cptr, unsafe.Pointer(&data), C.uint8_t(length)))
}

func (r *R) StartWrite(data []byte, length uint8) {
	C.rf24_startWrite(r.cptr, unsafe.Pointer(&data), C.uint8_t(length))
}

func (r *R) WriteAckPayload(pipe uint8, data []byte, length uint8) {
	C.rf24_writeAckPayload(r.cptr, C.uint8_t(pipe), unsafe.Pointer(&data), C.uint8_t(length))
}

func (r *R) Available() bool {
	return gobool(C.rf24_available(r.cptr))
}

// TODO: create Pipe type, make this a func on Pipe
func (r *R) AvailablePipe() (bool, uint8) {
	var pipe C.uint8_t
	available := gobool(C.rf24_available_pipe(r.cptr, &pipe))
	return available, uint8(pipe)
}

func (r *R) IsAckPayloadAvailable() bool {
	return gobool(C.rf24_isAckPayloadAvailable(r.cptr))
}


func (r *R) Read(length uint8) ([]byte, bool) {
	ok := gobool(C.rf24_read(r.cptr, unsafe.Pointer(&r.buffer[0]), C.uint8_t(length)))
	return r.buffer[:length],ok
}

func (r *R) OpenWritingPipe(address uint64) {
	C.rf24_openWritingPipe(r.cptr, C.uint64_t(address))
}

func (r *R) OpenReadingPipe(pipe uint8, address uint64) {
	C.rf24_openReadingPipe(r.cptr, C.uint8_t(pipe), C.uint64_t(address))
}

func (r *R) SetRetries(delay uint8, count uint8) {
	C.rf24_setRetries(r.cptr, C.uint8_t(delay), C.uint8_t(count))
}

func (r *R) SetChannel(channel uint8) {
	C.rf24_setChannel(r.cptr, C.uint8_t(channel))
}

func (r *R) SetPayloadSize(size uint8) {
	C.rf24_setPayloadSize(r.cptr, C.uint8_t(size))
}

func (r *R) GetPayloadSize() uint8 {
	return uint8(C.rf24_getPayloadSize(r.cptr))
}

func (r *R) GetDynamicPayloadSize() uint8 {
	return uint8(C.rf24_getDynamicPayloadSize(r.cptr))
}

func (r *R) EnableAckPayload() {
	C.rf24_enableAckPayload(r.cptr)
}

func (r *R) EnableDynamicPayloads() {
	C.rf24_enableDynamicPayloads(r.cptr)
}

func (r *R) IsPVariant() bool {
	return gobool(C.rf24_isPVariant(r.cptr))
}

func (r *R) SetAutoAck(enable bool) {
	C.rf24_setAutoAck(r.cptr, cbool(enable))
}

func (r *R) SetAutoAckPipe(pipe uint8, enable bool) {
	C.rf24_setAutoAck_pipe(r.cptr, C.uint8_t(pipe), cbool(enable))
}



// Is there any way to use the rf24_pa_dbm_e enum type directly
func (r *R) SetPALevel(level PA_DBM) {
	C.rf24_setPALevel(r.cptr, C.rf24_pa_dbm_val(level))
}

func (r *R) GetPALevel() PA_DBM {
	return PA_DBM(C.rf24_getPALevel(r.cptr))
}



func (r *R) SetDataRate(rate DATARATE) {
	C.rf24_setDataRate(r.cptr, C.rf24_datarate_val(rate))
}

func (r *R) GetDataRate() DATARATE {
	return DATARATE(C.rf24_getDataRate(r.cptr))
}



func (r *R) SetCRCLength(length CrcLength) {
	C.rf24_setCRCLength(r.cptr, C.rf24_crclength_val(length))
}

func (r *R) GetCRCLength() CrcLength {
	return CrcLength(C.rf24_getCRCLength(r.cptr))
}

func (r *R) DisableCRC() {
	C.rf24_disableCRC(r.cptr)
}

// TODO: String() method would be great
func (r *R) PrintDetails() {
	C.rf24_printDetails(r.cptr)
}

func (r *R) PowerDown() {
	C.rf24_powerDown(r.cptr)
}

func (r *R) PowerUp() {
	C.rf24_powerUp(r.cptr)
}

func (r *R) WhatHappened() (tx_ok, tx_fail, rx_ready bool) {
	var out_tx_ok, out_tx_fail, out_rx_ready C.cbool
	C.rf24_whatHappened(r.cptr, &out_tx_ok, &out_tx_fail, &out_rx_ready)
	tx_ok, tx_fail, rx_ready = gobool(out_tx_ok), gobool(out_tx_fail), gobool(out_rx_ready)
	return
}

func (r *R) TestCarrier() bool {
	return gobool(C.rf24_testCarrier(r.cptr))
}

func (r *R) TestRPD() bool {
	return gobool(C.rf24_testRPD(r.cptr))
}

// TODO: move out to util.go
func gobool(b C.cbool) bool {
	if b == C.cbool(0) {
		return false
	}

	return true
}

func cbool(b bool) C.cbool {
	if b {
		return C.cbool(1)
	}

	return C.cbool(0)
}
*/

/*
func main() {
	var pipe uint64 = 0xF0F0F0F0E1

	r := New("/dev/spidev0.0", 8000000, 25)
	defer r.Delete()

	r.Begin()
	r.SetRetries(15, 15)
	r.SetAutoAck(true)
	r.OpenReadingPipe(1, pipe)
	r.StartListening()
	r.PrintDetails()

	for {
		if r.Available() {
			data, _ := r.Read(4)
//			fmt.Printf("data: %v\n", data)
			payload := binary.LittleEndian.Uint32(data)
			fmt.Printf("Received %v\n", payload)

		} else {
			//
		}
	}
}
*/
