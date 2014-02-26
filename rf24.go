/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */

package gorf24

import (
	"time"

	"github.com/galaktor/gorf24/cmd"
	"github.com/galaktor/gorf24/gpio"
	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/spi"

	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/xaddr"
	/* registers */
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
	"github.com/galaktor/gorf24/reg/rxaddr"
	"github.com/galaktor/gorf24/reg/rxpw"
	"github.com/galaktor/gorf24/reg/status"
	"github.com/galaktor/gorf24/reg/txaddr"
	"github.com/galaktor/gorf24/reg/txobs"
)

const RF24_PAYLOAD_SIZE = 32

var EMPTY = make([]byte,0,0) // used when sending 'nothing'

type R struct {
	/*** I/O ***/
	spi *spi.SPI
	ce  *gpio.Pin
	csn *gpio.Pin

	/*** REGISTERS ***/
	/* in order of appearance in spec for now */
	config   *config.C
	autoAck  *autoack.AA // ?new on nRF24L01+
	enRxAddr *enrxaddr.E
	addrWid  *addrw.AW
	retrans  *retrans.R
	rfchan   *rfchan.R
	rfsetup  *rfsetup.R
	status   *status.S
	trans    *txobs.O
	rpd      *rpd.R // was 'CD' before nRF24L01+
	rxAddrP0 *xaddr.Full
	rxAddrP1 *xaddr.Full
	rxAddrP2 *xaddr.Partial
	rxAddrP3 *xaddr.Partial
	rxAddrP4 *xaddr.Partial
	rxAddrP5 *xaddr.Partial
	txAddr   *xaddr.Full
	rxPwP0   *rxpw.W
	rxPwP1   *rxpw.W
	rxPwP2   *rxpw.W
	rxPwP3   *rxpw.W
	rxPwP4   *rxpw.W
	rxPwP5   *rxpw.W
	fifo     *fifo.F
	// ACK_PLD set by command
	// TX_PLD set by command
	// RX_PLD set by command
	dynpd *dynpd.DP
	feat  *feature.F
}

func New(spidevice string, spispeed uint32, cepin, csnpin uint8) (r *R, err error) {
	r = &R{}
	r.config = config.New()
	r.autoAck = autoack.New()
	r.enRxAddr = enrxaddr.New()
	r.addrWid = addrw.New()
	r.retrans = retrans.New()
	r.rfchan = rfchan.New()
	r.rfsetup = rfsetup.New()
	r.status = status.New()
	r.trans = txobs.New()
	r.rpd = rpd.New()
	r.rxAddrP0 = rxaddr.NewFull(pipe.P0, xaddr.NewFromI(0))
	r.rxAddrP1 = rxaddr.NewFull(pipe.P1, xaddr.NewFromI(0))
	r.rxAddrP2 = rxaddr.NewPartial(pipe.P2, r.rxAddrP1, 0)
	r.rxAddrP3 = rxaddr.NewPartial(pipe.P3, r.rxAddrP1, 0)
	r.rxAddrP4 = rxaddr.NewPartial(pipe.P4, r.rxAddrP1, 0)
	r.rxAddrP5 = rxaddr.NewPartial(pipe.P5, r.rxAddrP1, 0)
	r.txAddr = txaddr.New(xaddr.NewFromI(0))
	r.rxPwP0 = rxpw.New(pipe.P0)
	r.rxPwP1 = rxpw.New(pipe.P1)
	r.rxPwP2 = rxpw.New(pipe.P2)
	r.rxPwP3 = rxpw.New(pipe.P3)
	r.rxPwP4 = rxpw.New(pipe.P4)
	r.rxPwP5 = rxpw.New(pipe.P5)
	r.fifo = fifo.New()
	r.dynpd = dynpd.New()
	r.feat = feature.New()

	r.spi, err = spi.New(spidevice, 0, 8, spi.SPD_02MHz)
	if err != nil {
		return
	}

	r.ce, err = gpio.Open(cepin, gpio.OUT)
	if err != nil {
		return
	}

	r.csn, err = gpio.Open(csnpin, gpio.OUT)
	if err != nil {
		return
	}

	r.ce.SetLow()
	r.csn.SetHigh()

	// ** FROM RF24.cpp **
	// Must allow the radio time to settle else configuration bits will not necessarily stick.
	// This is actually only required following power up but some settling time also appears to
	// be required after resets too. For full coverage, we'll always assume the worst.
	// Enabling 16b CRC is by far the most obvious case if the wrong timing is used - or skipped.
	// Technically we require 4.5ms + 14us as a worst case. We'll just call it 5ms for good measure.
	// WARNING: Delay is based on P-variant whereby non-P *may* require different timing.
	<-time.After(5 * time.Millisecond)

	return
}

func (r *R) Status() *status.S {
	return r.status
}

/*
Write command and status registers. AAAAA = 5
bit Register Map Address
Executable in power down or standby modes
only.
*/
func (rf *R) write(r reg.Register) error {
	rf.csn.SetLow()
	defer rf.csn.SetHigh()

	err := rf.sendCommand(cmd.W_REGISTER(r))
	if err != nil {
		return err
	}

	// write register bytes to spi, ignore what comes back
	_, err = r.WriteTo(rf.spi)
	if err != nil {
		return err
	}

	return nil
}

/*
Read command and status registers. AAAAA =
5 bit Register Map Address
*/
func (rf *R) read(r reg.Register) error {
	rf.csn.SetLow()
	defer rf.csn.SetHigh()

	err := rf.sendCommand(cmd.R_REGISTER(r))
	if err != nil {
		return err
	}

	// read register bytes from spi, send gibberish - only receive
	_, err = r.WriteTo(rf.spi)
	if err != nil {
		return err
	}

	return nil
}

// WARNING: expects CSN to be low already
// strictly just for use in func managing SPI already
// TODO? check for (cached) CSN state without actual read?
func (r *R) sendCommand(c cmd.C) error {
	s, err := r.spi.Transfer(c.Byte())
	if err != nil {
		return err
	}
	r.status.Set(s)
	return nil
}

func (r *R) writeCmd(c cmd.C, buf []byte) error {
	r.csn.SetLow()
	defer r.csn.SetHigh()

	r.sendCommand(c)
	_,err := r.spi.Write(buf)
	return err
}

func (r *R) readCmd(c cmd.C, buf []byte) error {
	r.csn.SetLow()
	defer r.csn.SetHigh()

	r.sendCommand(c)
	_,err := r.spi.Write(buf)
	return err
}


// TODO: check for payload widths? set to Tx/Rx modes when called?

// TODO: have different interface for PRX and PTX radios?

/*
Read RX-payload: 1 – 32 bytes. A read operation
always starts at byte 0. Payload is deleted from
FIFO after it is read. Used in RX mode
*/
func (r *R) readPayload(buf []byte) error {
	return r.readCmd(cmd.R_RX_PAYLOAD, buf)
}



/*
Write TX-payload: 1 – 32 bytes. A write operation
always starts at byte 0 used in TX payload.
*/
func (r *R) writePayload(buf []byte) error {
	return r.writeCmd(cmd.W_TX_PAYLOAD, buf)
}



// ?TODO: enum for modes, MD_RX and MD_TX?

/*
Flush TX FIFO, used in TX mode
*/
func (r *R) flushTx() error {
	return r.writeCmd(cmd.FLUSH_TX, EMPTY)
}

/*
Flush RX FIFO, used in RX mode
Should not be executed during transmission of
acknowledge, that is, acknowledge package will
not be completed.
*/
func (r *R) flushRx() error {
	// ?TODO: enforce/check mode?
	// from spec:
	//   Should not be executed during transmission of
	//   acknowledge, that is, acknowledge package will
	//   not be completed.
	return r.writeCmd(cmd.FLUSH_RX, EMPTY)
}

/*
No Operation. Might be used to read the STATUS
register
*/
func (r *R) refreshStatus() (*status.S, error) {
	err := r.writeCmd(cmd.NOP, EMPTY)
	return r.status, err
}

/*
This write command followed by data 0x73 acti-
vates the following features:
• R_RX_PL_WID
• W_ACK_PAYLOAD
• W_TX_PAYLOAD_NOACK
A new ACTIVATE command with the same data
deactivates them again. This is executable in
power down or stand by modes only.
The R_RX_PL_WID, W_ACK_PAYLOAD, and
W_TX_PAYLOAD_NOACK features registers are
initially in a deactivated state; a write has no
effect, a read only results in zeros on MISO. To
activate these registers, use the ACTIVATE com-
mand followed by data 0x73. Then they can be
accessed as any other register in nRF24L01. Use
the same command and data to deactivate the
registers again.
*/
func (r *R) toggleActivate() error {
	// TODO: keep activated bool state, make de/activate() funcs
	return r.writeCmd(cmd.ACTIVATE, []byte{0x73})
}
