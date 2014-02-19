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
