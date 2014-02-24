/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"errors"
//	"fmt"
//	"bytes"
	"io"

	"github.com/galaktor/gorf24/reg/addr"
)

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
	return 0,errors.New("not implemented")
}

/* Reader now implemented in r.reader (io.Reader)
// IO.WRITER
func (r *R) Write(p []byte) (n int, err error) {
	if len(p) == 1 {
		r.Set(p[0])
		return 1,nil
	}

	return 0,errors.New(fmt.Sprintf("cannot write %v bytes. register can only store exactly 1 byte", len(p)))
}

// IO.READER
func (r *R) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0,errors.New("cannot read into zero-length buffer. Register requires 1 byte.")
	}

	p[0] = r.flags
	return 1,nil
}
*/







