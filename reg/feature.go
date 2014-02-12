package reg

import (
	"github.com/galaktor/gorf24/reg/addr"
)

type Feature struct {
	R
}

func NewFeature(flags byte) *Feature {
	return &Feature{R{a: addr.FEATURE, flags: flags}}
}
