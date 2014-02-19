package xaddr

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* has the same 5byte address as its 'root' except for
   the LSByte which differs. When asked will combine
   the root's first 4 MSBytes and append it's own single
   LSByte to form a 5byte address. */
type Partial struct {
	reg.R

	root *Full
}

func NewPartial(a addr.A, root *Full, lsb byte) *Partial {
	return &Partial{reg.New(a, lsb), root}
}

func (r *Partial) Get() A {
	// use New method to force truncate
	return NewA((r.root.Get().Byte() << 8) | uint64(r.R.Byte()))
}

func (r *Partial) Set(lsb byte) {
	r.R.Set(lsb)
}
