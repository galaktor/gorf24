package nRF24L01P

import (
	"github.com/galaktor/gorf24/reg/addrw"
	"github.com/galaktor/gorf24/reg/autoack"
	"github.com/galaktor/gorf24/reg/config"
	"github.com/galaktor/gorf24/reg/dynpd"
	"github.com/galaktor/gorf24/reg/enrxaddr"
	"github.com/galaktor/gorf24/reg/feature"
	"github.com/galaktor/gorf24/reg/fifo"
	"github.com/galaktor/gorf24/reg/retrans"
	"github.com/galaktor/gorf24/reg/rfchan"
	"github.com/galaktor/gorf24/reg/rfsetup"
	"github.com/galaktor/gorf24/reg/rpd"
	"github.com/galaktor/gorf24/reg/rxpw"
	"github.com/galaktor/gorf24/reg/status"
	"github.com/galaktor/gorf24/reg/txobs"
	"github.com/galaktor/gorf24/reg/xaddr"
)

type Radio struct {
}

func New() *Radio {
	return nil
}

/* REGISTERS */
// implement StandbyRegGetter
func (r *Radio) GetConfig() (config.C, error)        { return config.C{}, nil }
func (r *Radio) GetAutoAck() (autoack.AA, error)     { return autoack.AA{}, nil }
func (r *Radio) GetEnRxAddr() (enrxaddr.E, error)    { return enrxaddr.E{}, nil }
func (r *Radio) GetAddrWid() (addrw.AW, error)       { return addrw.AW{}, nil }
func (r *Radio) GetRetrans() (retrans.R, error)      { return retrans.R{}, nil }
func (r *Radio) GetRfchan() (rfchan.R, error)        { return rfchan.R{}, nil }
func (r *Radio) GetRfsetup() (rfsetup.R, error)      { return rfsetup.R{}, nil }
func (r *Radio) GetStatus() (status.S, error)        { return status.S{}, nil }
func (r *Radio) GetTrans() (txobs.O, error)          { return txobs.O{}, nil }
func (r *Radio) GetRxAddrP0() (xaddr.Full, error)    { return xaddr.Full{}, nil }
func (r *Radio) GetRxAddrP1() (xaddr.Full, error)    { return xaddr.Full{}, nil }
func (r *Radio) GetRxAddrP2() (xaddr.Partial, error) { return xaddr.Partial{}, nil }
func (r *Radio) GetRxAddrP3() (xaddr.Partial, error) { return xaddr.Partial{}, nil }
func (r *Radio) GetRxAddrP4() (xaddr.Partial, error) { return xaddr.Partial{}, nil }
func (r *Radio) GetRxAddrP5() (xaddr.Partial, error) { return xaddr.Partial{}, nil }
func (r *Radio) GetTxAddr() (xaddr.Full, error)      { return xaddr.Full{}, nil }
func (r *Radio) GetRxPwP0() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetRxPwP1() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetRxPwP2() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetRxPwP3() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetRxPwP4() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetRxPwP5() (rxpw.W, error)          { return rxpw.W{}, nil }
func (r *Radio) GetFifo() (fifo.F, error)            { return fifo.F{}, nil }
func (r *Radio) GetDynpd() (dynpd.DP, error)         { return dynpd.DP{}, nil }
func (r *Radio) GetFeat() (feature.F, error)         { return feature.F{}, nil }

// implement StandbyRegSetter
func (r *Radio) SetConfig(f func(c *config.C)) error        { return nil }
func (r *Radio) SetAutoAck(f func(a *autoack.AA)) error     { return nil }
func (r *Radio) SetEnRxAddr(f func(e *enrxaddr.E)) error    { return nil }
func (r *Radio) SetAddrWid(f func(a *addrw.AW)) error       { return nil }
func (r *Radio) SetRetrans(f func(r *retrans.R)) error      { return nil }
func (r *Radio) SetRfchan(f func(r *rfchan.R)) error        { return nil }
func (r *Radio) SetRfsetup(f func(r *rfsetup.R)) error      { return nil }
func (r *Radio) SetStatus(f func(s *status.S)) error        { return nil }
func (r *Radio) SetTrans(f func(t *txobs.O)) error          { return nil }
func (r *Radio) SetRpd(f func(r *rpd.R)) error              { return nil }
func (r *Radio) SetRxAddrP0(f func(a *xaddr.Full)) error    { return nil }
func (r *Radio) SetRxAddrP1(f func(a *xaddr.Full)) error    { return nil }
func (r *Radio) SetRxAddrP2(f func(a *xaddr.Partial)) error { return nil }
func (r *Radio) SetRxAddrP3(f func(a *xaddr.Partial)) error { return nil }
func (r *Radio) SetRxAddrP4(f func(a *xaddr.Partial)) error { return nil }
func (r *Radio) SetRxAddrP5(f func(a *xaddr.Partial)) error { return nil }
func (r *Radio) SetTxAddr(f func(a *xaddr.Full)) error      { return nil }
func (r *Radio) SetRxPwP0(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetRxPwP1(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetRxPwP2(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetRxPwP3(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetRxPwP4(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetRxPwP5(f func(w *rxpw.W)) error          { return nil }
func (r *Radio) SetFifo(f func(*fifo.F)) error              { return nil }
func (r *Radio) SetDynpd(f func(d *dynpd.DP)) error         { return nil }
func (r *Radio) SetFeat(f func(f *feature.F)) error         { return nil }

// implement TxRegGetter

// implement RxRegGetter
func (r *Radio) GetRpd() rpd.R { return rpd.R{} }

/* MODES */
// implement ModeSwitcher
func (r *Radio) Standby() StandbyMode { return nil }
func (r *Radio) Tx() TxMode           { return nil }
func (r *Radio) Rx() RxMode           { return nil }

// implement StandbyMode

// implement TxMode

// implement RxMode
