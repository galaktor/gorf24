/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* a register with a one-byte address and one-byte data */
type R struct {
	a     addr.A
	flags byte
}

func New(a addr.A, flags byte) R {
	return R{a, flags}
}

func (r *R) Byte() byte {
	return r.flags
}

func (r *R) Set(flags byte) {
	r.flags = flags
}

func (r *R) Address() addr.A {
	return r.a
}
