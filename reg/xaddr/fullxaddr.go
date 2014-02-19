/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* stores a 5byte address.
   note that the R.flags field is left zero */
type Full struct {
	reg.R

	flags A
}

func NewFull(a addr.A, flags A) *Full {
	return &Full{reg.New(a, 0), flags}
}

func (r *Full) Get() A {
	return r.flags
}

func (r *Full) Set(a A) {
	r.flags = a
}
