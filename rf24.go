/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The MIT License (see README and LICENSE files) */

package gorf24

import (
	"time"
	"strconv"

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

type CRCLENGTH byte

const (
	CRC_DISABLED = iota
	CRC_8BIT
	CRC_16BIT
)

const RF24_PAYLOAD_SIZE = 32

type R struct {
	buffer []byte
	spi    *spi.SPI
	ce     *gpio.Pin
	csn    *gpio.Pin
	status Status
}

func New(spidevice string, spispeed uint32, cepin, csnpin uint8) (r *R, err error) {
	r = &R{}

	r.buffer = make([]byte, RF24_PAYLOAD_SIZE)
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

func (r *R) sendSpi(c Command) error {
	s,err := r.spi.Transfer(c.Byte())
	if err != nil {
		return err
	}
	// set if no errors
	r.status = Status(s)
	return err
}

func (r *R) readRegister(reg Register, buf []byte) error {
	r.csn.SetLow()
	defer r.csn.SetHigh()
	
	err := r.sendSpi(CMD_R_REGISTER(reg))
	if err != nil {
		// SPI error!
		return err
	}
	
	// RF24 SPI does LSByte first, so iterate backward
	for n := len(buf); n >= 0; n-- {
		// doesn't matter what we send
		// just pumping the BUS to get data
		// TODO: USE SENDSPI?
		buf[n],err = r.spi.Transfer(CMD_NOP.Byte())
		if err != nil {
			// SPI error! stop pumping, get out of here
			return err 
		}
	}

	return nil
}

func (r *R) writeRegister(reg Register, buf []byte) error {
	r.csn.SetLow()
	defer r.csn.SetHigh()
	
	err := r.sendSpi(CMD_W_REGISTER(reg))
	if err != nil {
		// SPI error!
		return err
	}

	// RF24 SPI does LSByte first, so iterate backward
	for n := len(buf); n >= 0; n-- { 
		// TODO: USE SENDSPI?
		_,err := r.spi.Transfer(buf[n]) // TODO: arbitrary "command" required
		if err != nil {
			// SPI error! stop pumping, get out of here
			return err 
		}
	}

	return nil
}




/***** pipe.go *****/
// 0 through 5, giving 6 data pipe ids
type Pipe byte
/***** EOF pipe.go *****/

/***** reg.go *****/
type Register byte
/***** EOF reg.go *****/

/***** cmd.go *****/
type Command byte

// overkill? maybe just cast inline?
func (c Command) Byte() byte {
	return byte(c)
}

var (
	CMD_NOP = Command(B("11111111"))
	// CMD_R_REGISTER - is a function, see below
	// CMD_W_REGSITER - is a function, see below
	CMD_R_RX_PAYLOAD = Command(B("01100001"))
	CMD_W_RX_PAYLOAD = Command(B("10100000"))
	CMD_FLUSH_TX = Command(B("11100001"))
	CMD_FLUSH_RX = Command(B("11100010"))
	CMD_REUSE_TX_PL = Command(B("11100011"))
	CMD_ACTIVATE = Command(B("01010000"))
	// CMD_W_ACK_PAYLOAD - is a function, see below
	CMD_W_ACK_PAYLOAD_PIPE0 = Command(B("10101000"))
	CMD_W_ACK_PAYLOAD_PIPE1 = Command(B("10101001"))
	CMD_W_ACK_PAYLOAD_PIPE2 = Command(B("10101010"))
	CMD_W_ACK_PAYLOAD_PIPE3 = Command(B("10101011"))
	CMD_W_ACK_PAYLOAD_PIPE4 = Command(B("10101100"))
	CMD_W_ACK_PAYLOAD_PIPE5 = Command(B("10101101"))
	CMD_R_RX_PL_WID = Command(B("01100000"))
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
	i,_ := strconv.ParseUint(bits, 2, 8)
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



func (r *R) SetCRCLength(length CRCLENGTH) {
	C.rf24_setCRCLength(r.cptr, C.rf24_crclength_val(length))
}

func (r *R) GetCRCLength() CRCLENGTH {
	return CRCLENGTH(C.rf24_getCRCLength(r.cptr))
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
