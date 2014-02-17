package reg

import (
	"errors"
	"fmt"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/pipe"
)

// 

const PW_MAX_WIDTH uint8 = 32

type RxPayloadWidth struct {
	R
}

func NewRxPayloadWidth(p pipe.P, flags byte) *RxPayloadWidth {
	return &RxPayloadWidth{R{a: addr.RX_PW(p), flags: flags}}
}

/* RX_PWx (bits 5:0)
   bits 7:6 reserved, must be '00' */
func (r *RxPayloadWidth) Set(w uint8) error {
	if w > PW_MAX_WIDTH {
		return errors.New(fmt.Sprintf("value out of legal range: %v. allowed values from 0 - %v", w, PW_MAX_WIDTH))
	}
	
	r.flags = (r.flags & 0) | w
	return nil
}
func (r *RxPayloadWidth) Get() uint8 {
	if r.flags > PW_MAX_WIDTH {
		return PW_MAX_WIDTH
	}

	return r.flags
}
