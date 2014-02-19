/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package spi

/*
#include <sys/ioctl.h>
#include <linux/spi/spidev.h>

unsigned long int spi_ioc_message(int i) {
  return SPI_IOC_MESSAGE(i);
}

// cgo cannot handle '...' parameter, so need to wrap ioctl
int Ioctl(int fd, unsigned long int request, void* out) {
  return ioctl(fd, request, out);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

type SpiSpeed uint32

const (
	SPD_500KHz SpiSpeed =   500000
	SPD_02MHz  SpiSpeed =  2000000
	SPD_08MHz  SpiSpeed =  8000000
	SPD_16MHz  SpiSpeed = 16000000
)

type SPI struct {
	bitsPerWord byte
	speedHz     SpiSpeed
	mode        byte
	device      *os.File

	rxBuf       byte
	tr *C.struct_spi_ioc_transfer
}

func New(d string, m byte, b byte, s SpiSpeed) (*SPI, error) {
	// try simple open; might need os.OpenFile, or syscall.Open
	f, err := os.Open(d)
	if err != nil {
		return nil, err
	}
	spi := &SPI{device: f, tr: &C.struct_spi_ioc_transfer{}}
	spi.tr.rx_buf = C.__u64(uintptr(unsafe.Pointer(&spi.rxBuf)))

	spi.setMode(m)
	spi.setBitsPerWord(b)
	spi.setSpeed(s)

	return spi,nil
}

func (s *SPI) Delete() {
	s.device.Close()
}

func (s *SPI) getDevice() string {
	return s.device.Name()
}

func (s *SPI) setMode(m byte) error {
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_RD_MODE, unsafe.Pointer(&m))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI write mode on %s to %v", s.getDevice(), m))
	}
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_WR_MODE, unsafe.Pointer(&m))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI read mode on %s to %v", s.getDevice(), m))
	}

	s.mode = m
	return nil
}

func (s *SPI) setBitsPerWord(b byte) error {
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_RD_BITS_PER_WORD, unsafe.Pointer(&b))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI read bits on %s to %v", s.getDevice(), b))
	}
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_WR_BITS_PER_WORD, unsafe.Pointer(&b))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI write bits on %s to %v", s.getDevice(), b))
	}
	s.bitsPerWord = b
	return nil
}

func (s *SPI) setSpeed(spd SpiSpeed) error {
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_RD_MAX_SPEED_HZ, unsafe.Pointer(&spd))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI read speed on %s to %v", s.getDevice(), spd))
	}
	if int(C.Ioctl(C.int(s.device.Fd()), C.SPI_IOC_WR_MAX_SPEED_HZ, unsafe.Pointer(&spd))) == -1 {
		return errors.New(fmt.Sprintf("error setting SPI write speed on %s to %v", s.getDevice(), spd))
	}
	s.speedHz = spd
	return nil
}

func (s *SPI) Transfer(tx byte) (rx byte, err error) {
	// address -> pointer -> unsigned pointer -> cast to unsigned long
	s.tr.tx_buf = C.__u64(uintptr(unsafe.Pointer(&tx)))
	s.tr.len = 1
	
	/* temp  overrides
	s.tr.speed_hz = s.speed
	s.tr.bits_per_word = s.bitsPerWord */

	if int(C.Ioctl(C.int(s.device.Fd()), C.spi_ioc_message(1), unsafe.Pointer(&s.tr))) == -1 {
		return 0,errors.New(fmt.Sprintf("error during transfer on %v", s.getDevice()))
	}

	return s.rxBuf,nil
}

/*

func (s *SPI) Write(p []byte) (n int, err error) {
	
}

func (s *SPI) Read(p []byte) (n int, err error) {

}

*/








