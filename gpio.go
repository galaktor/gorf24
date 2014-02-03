package gpio
/*
#include <stdio.h>
*/
import "C"

import (
	"os"
	"fmt"
)

// see: https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio

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

func Open(port int, d Direction) error {
	// TODO: portString = "%d",port
	// use simple string concat or something below instead of fmt?
	// BENCHMARKS

	// OPEN PIN
	f,err := os.Open("/sys/class/gpio/export")
	if err != nil {
		return err
	}
	_,err = f.WriteString(fmt.Sprintf("%d\n", port))
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
	case IN:  f.WriteString("in\n")
	case OUT: f.WriteString("out\n")
	}
	f.Close()

	return nil
}

func Close(port int) error {
	// CLOSE PIN
	f,err := os.Open("/sys/class/gpio/unexport")
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("%d\n", port))
	f.Close()
	
	return nil
}

func Read(port int) (int,error) {
	f,err := os.Open(fmt.Sprintf("/sys/class/gpio/gpio%d/value", port))
	if err != nil {
		return 0,err
	}
	defer f.Close()
	var out int
	_,err = fmt.Fscanf(f, "%d", &out)
	if err != nil {
		return 0,err
	}
	
	return out,nil
}

func Write(port, value PinState) error {
	// WRITE VALUE TO PIN
	f,err := os.Open(fmt.Sprintf("/sys/class/gpio/gpio%d/value", port))
	if err != nil {
		return err
	}
	switch value {
	case LOW: f.WriteString("0\n")
	case HIGH: f.WriteString("1\n")
	}
	f.Close()

	return nil
}






