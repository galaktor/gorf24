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














