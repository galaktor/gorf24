package spi

import (
	"bytes"
)

type Fake struct {
	buf *bytes.Buffer
}

func NewFake(flags byte) *Fake {
	buf := bytes.Buffer{}
	buf.WriteByte(flags)
	return &Fake{&buf}
}

// implements interface: Pumper
func (f *Fake) Pump(tx byte) (rx byte, err error) {
	f.buf.WriteByte(tx)
	return 0,nil
}

// implements interface: io.Writer
func (f *Fake) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

// implements interface: io.Reader
func (f *Fake) Read(p []byte) (n int, err error) {
	return f.buf.Read(p)
}





