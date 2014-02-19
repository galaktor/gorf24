/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
// see: https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio
package gpio

import (
	"fmt"
	"os"
)

type Direction byte

const (
	IN  Direction = 0
	OUT Direction = 1
)

// TODO: const( PINS ? )

type PinState byte

const (
	LOW  PinState = 0
	HIGH PinState = 1
)

/*
var p := gpio.Open(gpio.RPI_11) // provide alias for SPI etc?
defer p.Close()
switch p.Read() {
case gpio.HIGH: p.SetLow()
case gpio.LOW:  p.SetHigh()
}
*/

type Pin struct {
	pinNr         byte
	direction     Direction
	directionFile string
	valueFile     string
}

func (p *Pin) Nr() byte {
	return p.pinNr
}

func Open(pin byte, d Direction) (*Pin, error) {
	p := &Pin{pinNr: pin,
		direction:     d,
		directionFile: fmt.Sprintf("/sys/class/gpio/gpio%d/direction", pin),
		valueFile:     fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)}
	err := p.open()
	if err != nil {
		return p, err
	}

	err = p.setDirection(d)
	return p, err
}

func (p *Pin) open() error {
	f, err := os.Open("/sys/class/gpio/export")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d\n", p.pinNr))
	return err
}

func (p *Pin) setDirection(d Direction) error {
	f, err := os.Open(p.directionFile)
	if err != nil {
		return err
	}
	switch d {
	case IN:
		f.WriteString("in\n")
	case OUT:
		f.WriteString("out\n")
	}
	err = f.Close()
	return err
}

func (p *Pin) Close() error {
	f, err := os.Open("/sys/class/gpio/unexport")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d\n", p.pinNr))
	return err
}

func (p *Pin) Read() (PinState, error) {
	f, err := os.Open(p.valueFile)
	if err != nil {
		return LOW, err
	}
	defer f.Close()
	var out PinState
	_, err = fmt.Fscanf(f, "%d", &out)
	if err != nil {
		return LOW, err
	}
	return out, nil
}

func (p *Pin) SetLow() error {
	return p.set(LOW)
}

func (p *Pin) SetHigh() error {
	return p.set(HIGH)
}

func (p *Pin) set(s PinState) error {
	f, err := os.Open(p.valueFile)
	if err != nil {
		return err
	}
	switch s {
	case LOW:
		f.WriteString("0\n")
	case HIGH:
		f.WriteString("1\n")
	}
	err = f.Close()
	return err
}

/* OLD CODE

func Open(port int, d Direction) error {
	// TODO: portString = "%d",port
	// use simple string concat or something below instead of fmt?
	// BENCHMARKS

	// OPEN PIN
	f, err := os.Open("/sys/class/gpio/export")
	if err != nil {
		return err
	}
	_, err = f.WriteString(fmt.Sprintf("%d\n", port))
	if err != nil {
		return err
	}
	f.Close()

	// SET PIN DIRECTION
	f, err = os.Open(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", port))
	if err != nil {
		return err
	}
	switch d {
	case IN:
		f.WriteString("in\n")
	case OUT:
		f.WriteString("out\n")
	}
	f.Close()

	return nil
}

func Close(port int) error {
	// CLOSE PIN
	f, err := os.Open("/sys/class/gpio/unexport")
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("%d\n", port))
	f.Close()

	return nil
}

func Read(port int) (int, error) {
	f, err := os.Open(fmt.Sprintf("/sys/class/gpio/gpio%d/value", port))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	var out int
	_, err = fmt.Fscanf(f, "%d", &out)
	if err != nil {
		return 0, err
	}

	return out, nil
}

func Write(port, value PinState) error {
	// WRITE VALUE TO PIN
	f, err := os.Open(fmt.Sprintf("/sys/class/gpio/gpio%d/value", port))
	if err != nil {
		return err
	}
	switch value {
	case LOW:
		f.WriteString("0\n")
	case HIGH:
		f.WriteString("1\n")
	}
	f.Close()

	return nil
}
*/
