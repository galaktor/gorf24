/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package reg

import (
	"errors"
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/util"
)

func TestNew_AddressFortyTwo_StoresThatAddress(t *testing.T) {
	expected := addr.A(42)
	r := New(expected, 0)

	actual := r.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestNew_MaskFortyTwo_StoresThaMask(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), expected)

	actual := r.Mask()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestNO_MASK_Is_AllOnes(t *testing.T) {
	expected := util.B("11111111")

	actual := NO_MASK

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestSet_FortyTwo_NoMask_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)

	r.Set(expected)

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestSet_FortyTwo_ZeroMask_StoresZero(t *testing.T) {
	expected := byte(0)
	r := New(addr.A(0), 0x00)

	r.Set(expected)

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestSet_AllOnes_MixedMask_StoresMaskedValue(t *testing.T) {
	expected := byte(0xAA)
	r := New(addr.A(0), 0xAA)

	r.Set(util.B("11111111"))

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestGet_FortyTwo_NoMask_ReturnsFortyTwo(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)
	r.Set(42)

	actual := r.Get()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWriteTo_WriterWithNoErrors_WritesOneByte_WithoutErrors(t *testing.T) {
	expected := int64(1)
	r := New(addr.A(0), NO_MASK)
	w := FakeRW{Passes()}

	actual, err := r.WriteTo(w)

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWriteTo_WriterWithNoErrors_WritesFlags(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)
	buf := make([]byte, 2)
	w := FakeRW{CopiesTo(buf)}
	r.Set(expected)

	r.WriteTo(w)

	actual := buf[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWriteTo_Masked_WritesMaskedFlags(t *testing.T) {
	expected := util.B("10101010")
	r := New(addr.A(0), util.B("10101010"))
	buf := make([]byte, 2)
	w := FakeRW{CopiesTo(buf)}
	r.Set(util.B("11111111"))

	r.WriteTo(w)

	actual := buf[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestReadFrom_EmptyReader_FlagsUnchanged(t *testing.T) {
	t.FailNow()
}

func TestReadFrom_EmptyReader_ReturnsZeroWithError(t *testing.T) {
	t.FailNow()
}

func TestReadFrom_ReaderNoErrors_MasksFirstByte(t *testing.T) {
	expected := util.B("10101010")
	r := New(addr.A(0), util.B("10101010"))
	buf := []byte{42}
	rd := FakeRW{CopiesFrom(buf)}

	r.ReadFrom(rd)

	actual := r.Get()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
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

func CopiesTo(out []byte) RwFunc {
	return func(p []byte) (int, error) {
		copy(out, p)
		return len(p), nil
	}
}

func CopiesFrom(in []byte) RwFunc {
	return func(p []byte) (int, error) {
		copy(p, in)
		return len(p), nil
	}
}

/*
func TestWrite_OneByte_NoMask_StoresThatByte(t *testing.T) {
	expected := byte(42)
	r := New(addr.A(0), NO_MASK)

	r.Write([]byte{expected})

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_OneByte_Masked_StoresMaskedByte(t *testing.T) {
	expected := byte(0xAA)
	r := New(addr.A(0), 0xAA)

	r.Write([]byte{util.B("11111111")})

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_OneByte_ReturnsNilError(t *testing.T) {
	r := New(addr.A(0), NO_MASK)

	_,err := r.Write([]byte{0})

	if err != nil {
		t.Errorf("unexpected error: '%v'", err)
	}
}

func TestWrite_OneByte_ReturnsOne(t *testing.T) {
	expected := 1
	r := New(addr.A(0), NO_MASK)

	actual,_ := r.Write([]byte{0})

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoBytes_DoesNotChangeFlags(t *testing.T) {
	expected := util.B("11111111")
	r := New(addr.A(0), NO_MASK)
	r.Set(expected)

	r.Write(make([]byte, 0))

	actual := r.flags[0]
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoByte_ReturnsError(t *testing.T) {
	expected := "cannot write 0 bytes. register can only store exactly 1 byte"
	r := New(addr.A(0), NO_MASK)

	_, err := r.Write(make([]byte, 0))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_TwoBytes_ReturnsError(t *testing.T) {
	expected := "cannot write 2 bytes. register can only store exactly 1 byte"
	r := New(addr.A(0), NO_MASK)

	_, err := r.Write(make([]byte, 2))

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_NoByte_ReturnsZero(t *testing.T) {
	expected := 0
	r := New(addr.A(0), NO_MASK)

	actual, _ := r.Write(make([]byte, 0))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestWrite_TwoBytes_ReturnsZero(t *testing.T) {
	expected := 0
	r := New(addr.A(0), NO_MASK)

	actual, _ := r.Write(make([]byte, 2))

	if actual != expected {
		t.Errorf("expected '%v' but found '%v' with reg '%v'", expected, actual, r)
	}
}

func TestRead_ZeroLengthSlice_ReturnsZeroAndError(t *testing.T) {
	expected := "cannot read into zero-length buffer. Register requires 1 byte."
	r := New(addr.A(0), NO_MASK)

	n,err := r.Read(make([]byte, 0))

	if n != 0 {
		t.Errorf("expected '%v' but foun '%v' with reg '%v'", 0, n, r)
	}

	if err == nil {
		t.Error("expected error but found nil")
		t.FailNow()
	}

	if err.Error() != expected {
		t.Errorf("expected '%v' but foun '%v' with reg '%v'", expected, err.Error(), r)
	}
}

*/
