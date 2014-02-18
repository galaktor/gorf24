package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

/* stores a 5byte address.
   note that the R.flags field is left zero */
type FullXAddress struct {
	R

	flags XAddress
}

func newFullXAddress(a addr.A, flags XAddress) *FullXAddress {
	return &FullXAddress{R{a, 0}, flags}
}

func (r *FullXAddress) Get() XAddress {
	return r.flags
}

func (r *FullXAddress) Set(a XAddress) {
	r.flags = a
}
