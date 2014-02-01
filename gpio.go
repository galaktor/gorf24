package gpio
/*
#include <stdio.h>
*/
import "C"

import (
	"os"
	"fmt"
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

func Open(port int, d Direction) error {
	// ? select pin ?
	f,err := os.Open("/sys/class/gpio/export")
	if err != nil {
		return err
	}
	_,err = f.WriteString(fmt.Sprintf("%d\n", port))
	if err != nil {
		return err
	}
	f.Close()

	// ? set direction ?
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

func Close(port int) {

}

func Read(port int) int {
	return -1
}

func Write(port, value int) {

}






