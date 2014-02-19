package txaddr

import (
	"testing"

	"github.com/galaktor/gorf24/reg/addr"
	"github.com/galaktor/gorf24/reg/xaddr"
)

func TestNew_RegisterAddress_IsTX_ADDR(t *testing.T) {
	a := New(xaddr.New(0))
	expected := addr.TX_ADDR

	actual := a.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with txaddr '%v'", expected, actual, a)
	}
}
