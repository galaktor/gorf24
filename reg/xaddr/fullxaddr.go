/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"io"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* stores a 5byte address.
   note that the R.flags field is left zero */
type Full struct {
	reg.R

	flags []byte // 5 bytes
}

func NewFull(a addr.A, flags A) *Full {
	return &Full{reg.New(a, 0x00), flags[:5]}
}

func (f *Full) Get() A {
	return NewFromB(f.flags)
}

func (f *Full) Set(a A) {
	f.flags = a[:5]
//	f.sliced = f.flags[:]
}

// io.ReaderFrom
func (f *Full) ReadFrom(r io.Reader) (n int64, err error) {
	n32,err := r.Read(f.flags)
	return int64(n32),nil
}










