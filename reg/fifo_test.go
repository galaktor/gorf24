package reg

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewFifoStatus_RegisterAddress_IsFIFO_STATUS(t *testing.T) {
	f := NewFifoStatus(0)
	expected := addr.FIFO_STATUS

	actual := f.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with fifostatus '%v'", expected, actual, f)
	}
}