package gorf24

/*
  #cgo LDFLAGS: -L./RF24_c
  #cgo LDFLAGS: -lrf24_c
  #cgo CFLAGS: -I./RF24_c
  #include "RF24_c.h"
  #include <stdio.h>
*/
import "C"

import (
//	"encoding/binary"
//	"fmt"
	"unsafe"
)

// TODO: more idiomatic:
//   error handling
//   Read/Write interfaces
//   possibly more conventional byte slice passing
//   more type safety?
//   channel for available data, like time.Tick? r.Available() <-chan []byte  or so?

type R struct {
	cptr C.RF24Handle
	buffer_size uint8
	buffer []byte
}

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

func New(spidevice string, spispeed uint32, cepin uint8) R {
	var r R
	r.cptr = C.new_rf24(C.CString(spidevice), C.uint32_t(spispeed), C.uint8_t(cepin))
	r.buffer = make([]byte, 128) // max payload length according to nrf24 spec

	return r
}

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

// TODO: move to named package
type PA_DBM byte

const (
	PA_MIN PA_DBM = iota
	PA_LOW
	PA_HIGH
	PA_MAX
	PA_ERROR // what is this for?
)

// Is there any way to use the rf24_pa_dbm_e enum type directly
func (r *R) SetPALevel(level PA_DBM) {
	C.rf24_setPALevel(r.cptr, C.rf24_pa_dbm_val(level))
}

func (r *R) GetPALevel() PA_DBM {
	return PA_DBM(C.rf24_getPALevel(r.cptr))
}

type DATARATE byte

const (
	RATE_1MBPS DATARATE = iota
	RATE_2MBPS
	RATE_250KBPS
)

func (r *R) SetDataRate(rate DATARATE) {
	C.rf24_setDataRate(r.cptr, C.rf24_datarate_val(rate))
}

func (r *R) GetDataRate() DATARATE {
	return DATARATE(C.rf24_getDataRate(r.cptr))
}

type CRCLENGTH byte

const (
	CRC_DISABLED = iota
	CRC_8BIT
	CRC_16BIT
)

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
