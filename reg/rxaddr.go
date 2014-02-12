package reg

import (
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/util"
)

type RxAddress [5]byte

func NewRxAddress(msb uint32, lsb byte) RxAddress {
	return util.ItoB5((uint64(msb) << 8) | uint64(lsb))
}

type RxAddresses struct {
	en *EnabledRxAddresses
	p0 *FullRxAddress // set RxAddrP0 equal to TX_ADDR when using AutoAck
	p1 *FullRxAddress
	p2 *PartialRxAddress
	p3 *PartialRxAddress
	p4 *PartialRxAddress
	p5 *PartialRxAddress
}

func NewRxAddresses() *RxAddresses {
	a := &RxAddresses{}
	a.en = NewEnabledRxAddresses(0)
	a.p0 = NewFullRxAddress(pipe.P0,[5]byte{})
	a.p1 = NewFullRxAddress(pipe.P1,[5]byte{})
	a.p2 = NewPartialRxAddress(pipe.P2, a.p1,0)
	a.p3 = NewPartialRxAddress(pipe.P3, a.p1,0)
	a.p4 = NewPartialRxAddress(pipe.P4, a.p1,0)
	a.p5 = NewPartialRxAddress(pipe.P5, a.p1,0)
	return a
}

func (r *RxAddresses) En() *EnabledRxAddresses {
	return r.en
}

func (r *RxAddresses) P0() *FullRxAddress {
	return r.p0
}

func (r *RxAddresses) P1() *FullRxAddress {
	return r.p1
}

func (r *RxAddresses) P2() *PartialRxAddress {
	return r.p2
}

func (r *RxAddresses) P3() *PartialRxAddress {
	return r.p3
}

func (r *RxAddresses) P4() *PartialRxAddress {
	return r.p4
}

func (r *RxAddresses) P5() *PartialRxAddress {
	return r.p5
}
