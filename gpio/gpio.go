/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
// see: https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio
package gpio

import (
	"fmt"
	"io"
	"os"
)

type Direction byte

const (
	IN  Direction = 0
	OUT Direction = 1
)

type PinState byte

const (
	LOW  PinState = 0
	HIGH PinState = 1
)

type Pin interface {
	io.Closer

	Nr() byte
	Direction() Direction

	Set(s PinState) error
	Get() (PinState, error)
}

/***** GPIO Open/Close *****/
func Open(pin byte, d Direction) (Pin, error) {
	p := &P{pinNr: pin,
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

func (p *P) open() error {
	f, err := os.Open("/sys/class/gpio/export")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d\n", p.pinNr))
	return err
}

func close(p *P) error {
	f, err := os.Open("/sys/class/gpio/unexport")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d\n", p.pinNr))
	return err
}

/***** Pin implementation *****/

/* P needs to open/close files ad-hoc
   alternative would be to hold file lock
   and store io.Reader/Writer instead
   but that would be "greedy" as it would
   prevent other apps from accessing the
   pins on the device */
type P struct {
	pinNr         byte
	direction     Direction
	directionFile string
	valueFile     string
}

func (p *P) Close() error {
	return close(p)
}

func (p *P) Nr() byte {
	return p.pinNr
}

func (p *P) Direction() Direction {
	return p.direction
}

func (p *P) setDirection(d Direction) error {
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

func (p *P) Get() (PinState, error) {
	f, err := os.Open(p.valueFile)
	// TODO: defer close?
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

func (p *P) Set(s PinState) error {
	f, err := os.Open(p.valueFile)
	// TODO: defer close?
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
