/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"testing"
	"errors"

	"github.com/galaktor/gorf24/reg/addr"
)

func TestGet_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := [5]byte{0, 0, 0, 0, 0}
	a := NewFull(addr.A(0), NewFromI(0))

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestGet_NonZero_ReturnsRightBits(t *testing.T) {
	expected := [5]byte{0xAF, 0xAF, 0xAF, 0xAF, 0xAF}
	a := NewFull(addr.A(0), NewFromI(0xAFAFAFAFAF))

	actual := a.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestSet_NonZero_StoresRightBits(t *testing.T) {
	expected := [5]byte{0xAF, 0xAF, 0xAF, 0xAF, 0xAF}
	a := NewFull(addr.A(0), NewFromI(0))

	a.Set(NewFromB5([5]byte{0xAF, 0xAF, 0xAF, 0xAF, 0xAF}))

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_Zeroes_SetsToZero(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}
	a := NewFull(addr.A(0), NewFromI(0))

	a.ReadFrom(FakeRW{From(0, 0, 0, 0, 0)})

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_MixedBytes_SetsToThoseBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE}
	a := NewFull(addr.A(0), NewFromI(0))

	a.ReadFrom(FakeRW{From(0xAA, 0xBB, 0xCC, 0xDD, 0xEE)})

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_MoreThan5Bytes_SetsToFiveMSBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE} // 5 bytes
	a := NewFull(addr.A(0), NewFromI(0))

	a.ReadFrom(FakeRW{From(0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF)}) // 6 bytes

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_LessThan5Bytes_SetsToMSBytes(t *testing.T) {
	expected := [5]byte{0xAA, 0xBB, 0xCC, 0x00, 0x00} // 5 bytes
	a := NewFull(addr.A(0), NewFromI(0))

	a.ReadFrom(FakeRW{From(0xAA, 0xBB, 0xCC)}) // 6 bytes

	actual := a.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_ReaderPasses_ReturnsReaderByteCount(t *testing.T) {
	expected := int64(42)
	a := NewFull(addr.A(0), NewFromI(0))

	actual,err := a.ReadFrom(FakeRW{PassesWith(42)})

	if err != nil {
		t.Errorf("expected nil error but found '%v'", err)
	}
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fulladdr '%v'", expected, actual, a)
	}
}

func TestReadFrom_ReaderError_ReturnsUnderlyingReaderError(t *testing.T) {
	expected := "some reader error occurred"
	a := NewFull(addr.A(0), NewFromI(0))
	
	n,err := a.ReadFrom(FakeRW{FailsWith(expected)})

	if n != 0 {
		t.Errorf("expected '%v' but found '%v' with fulladdr '%v'", 0, n, a)
	}

	if err == nil {
		t.Errorf("expected error '%v' but found nil", expected)
		t.FailNow()
	}
	
	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fulladr '%v'", expected, actual, a)
	}
}

func TestWriteTo_Zeroes_WritesZeroes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}
	a := NewFull(addr.A(0), NewFromI(0))
	actual := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	a.WriteTo(FakeRW{To(actual[:])})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' for fulladdr '%v'", expected, actual, a)
	}
}

func TestWriteTo_WriterPasses_ReturnsWriterByteCount(t *testing.T) {
	expected := int64(42)
	a := NewFull(addr.A(0), NewFromI(0))

	actual,err := a.WriteTo(FakeRW{PassesWith(42)})

	if err != nil {
		t.Errorf("expected nil error but found '%v'", err)
	}
	
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fulladdr '%v'", expected, actual, a)
	}
}

func TestWriteTo_WriterError_ReturnsUnderlyingWriterError(t *testing.T) {
	expected := "some reader error occurred"
	a := NewFull(addr.A(0), NewFromI(0))
	
	n,err := a.WriteTo(FakeRW{FailsWith(expected)})

	if n != 0 {
		t.Errorf("expected '%v' but found '%v' with fulladdr '%v'", 0, n, a)
	}

	if err == nil {
		t.Errorf("expected error '%v' but found nil", expected)
		t.FailNow()
	}
	
	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with fulladr '%v'", expected, actual, a)
	}
}

func TestWriteTo_FourByteTarget_WritesFourMSBytes(t *testing.T) {
	expected := [4]byte{0xAA, 0xBB, 0xCC, 0xDD}
	a := NewFull(addr.A(0), NewFromI(0xAABBCCDDEE))
	actual := [4]byte{}

	n,err := a.WriteTo(FakeRW{To(actual[:4]).Then(PassesWith(4))})

	if n != 4 {
		t.Errorf("expected '%v' but found '%v' with fulladdr '%v'", 4, n, a)
	}

	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}
	
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fulladdr '%v'", expected, actual, a)
	}
}

/***** FAKES *****/
type RwFunc func(p []byte) (n int, err error)

type FakeRW struct {
	cb RwFunc
}

func (f FakeRW) Read(p []byte) (int, error) {
	return f.cb(p)
}

func (f FakeRW) Write(p []byte) (int, error) {
	return f.cb(p)
}

func Fails() RwFunc {
	return FailsWith("this call is failing deliberately.")
}

func FailsWith(msg string) RwFunc {
	return func(p []byte) (int, error) { return 0, errors.New(msg) }
}

func Passes() RwFunc {
	return func(p []byte) (int, error) { return len(p), nil }
}

func PassesWith(n int) RwFunc {
	return func(p []byte) (int, error) { return n, nil }
}

func (first RwFunc) Then(then RwFunc) RwFunc {
	return func(p []byte) (n int, err error) {
		n, err = first(p)
		n, err = then(p)
		return
	}
}

func To(out []byte) RwFunc {
	return func(p []byte) (int, error) {
		copy(out, p)
		return len(p), nil
	}
}

func From(in ...byte) RwFunc {
	return func(p []byte) (int, error) {
		copy(p, in)
		return len(p), nil
	}
}














