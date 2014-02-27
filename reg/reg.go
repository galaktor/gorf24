/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"errors"
	"io"

	"github.com/galaktor/gorf24/reg/addr"
)

type Register interface {
	io.ReaderFrom
	io.WriterTo
	addr.IHaveAddress
}

/* a register with a one-byte address and one-byte data */
type R struct {
	a     addr.A
	mask  byte
	flags []byte
}

const NO_MASK byte = 0xFF // 11111111

func New(a addr.A, mask byte) R {
	return R{a, mask, make([]byte,1,1)}
}

func (r *R) Address() addr.A {
	return r.a
}

func (r *R) Mask() byte {
	return r.mask
}

func (r *R) Get() byte {
	return r.flags[0]
}

func (r *R) Set(flags byte) {
	r.flags[0] = flags & r.mask
}

// io.WriterTo
func (r *R) WriteTo(w io.Writer) (n int64, err error) {
	n32,err := w.Write(r.flags)
	n = int64(n32)
	return
}

// io.ReaderFrom
func (r *R) ReadFrom(rd io.Reader) (n int64, err error) {
	bkup := r.Get()
	n32,err := rd.Read(r.flags)
	n = int64(n32)

	if err != nil {
		r.Set(bkup) // roll back value
		return n,err
	}

	if n == 0 {
		return n,errors.New("could not read from empty Reader")
	}

	r.Set(r.flags[0]) // will mask

	return
}
