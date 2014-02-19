/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package pipe

// 0 through 5, giving 6 data pipe ids
type P byte

// TODO: consider setting pipe constants to bits as used in autoack etc regss
// can then just do "return flags & p == p"?

const (
	P0 P = iota
	P1
	P2
	P3
	P4
	P5
)