package reg

import (
	"testing"
	
//	"github.com/galaktor/gorf24/util"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewTransObserve_RegisterAddress_IsOBSERVE_TX(t *testing.T) {
	o := NewTransObserve(0)
	expected := addr.OBSERVE_TX

	actual := o.Address()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b' with transobs '%v'", expected, actual, o)
	}
}