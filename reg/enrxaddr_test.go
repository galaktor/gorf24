package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewRxAddresses_P0_SetToPipe0Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P0)

	a := NewRxAddresses()

	actual := a.P0().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P1_SetToPipe1Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P1)

	a := NewRxAddresses()

	actual := a.P1().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P2_SetToPipe2Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P2)

	a := NewRxAddresses()

	actual := a.P2().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P3_SetToPipe3Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P3)

	a := NewRxAddresses()

	actual := a.P3().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P4_SetToPipe4Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P4)

	a := NewRxAddresses()

	actual := a.P4().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P5_SetToPipe5Reg(t *testing.T) {
	expected := addr.RX_ADDR(pipe.P5)

	a := NewRxAddresses()

	actual := a.P5().Address()
	if actual != expected {
		t.Errorf("expected '%b' but found '%b' on rxaddresses '%v'", actual, expected, a)
	}
}

func TestNewRxAddresses_P2_MSByteReferenceP1Flags(t *testing.T) {
	a := NewRxAddresses()
	expected := &a.P1().flags[0]

	actual := &a.P2().msb[0]
	if actual != expected {
		t.Errorf("expected '%p' but found '%p' on rxaddresses '%v'", expected, actual, a)
	}
}

func TestNewRxAddresses_P3_MSByteReferenceP1Flags(t *testing.T) {
	a := NewRxAddresses()
	expected := &a.P1().flags[0]

	actual := &a.P3().msb[0]
	if actual != expected {
		t.Errorf("expected '%p' but found '%p' on rxaddresses '%v'", expected, actual, a)
	}
}

func TestNewRxAddresses_P4_MSByteReferenceP1Flags(t *testing.T) {
	a := NewRxAddresses()
	expected := &a.P1().flags[0]

	actual := &a.P4().msb[0]
	if actual != expected {
		t.Errorf("expected '%p' but found '%p' on rxaddresses '%v'", expected, actual, a)
	}
}

func TestNewRxAddresses_P5_MSByteReferenceP1Flags(t *testing.T) {
	a := NewRxAddresses()
	expected := &a.P1().flags[0]

	actual := &a.P5().msb[0]
	if actual != expected {
		t.Errorf("expected '%p' but found '%p' on rxaddresses '%v'", expected, actual, a)
	}
}

func TestNewRxAddresses_En_IsNotNil(t *testing.T) {
	a := NewRxAddresses()

	if a.En() == nil {
		t.Errorf("unexpected 'nil' on rxaddresses '%v'", a)
	}
}