package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* a register with a one-byte address and one-byte data */
type R struct {
	a     addr.A
	flags byte
}

func New(a addr.A) R {
	return R{a, byte(0)}
}

func (r *R) Byte() byte {
	return r.flags
}

func (r *R) Address() addr.A {
	return r.a
}
