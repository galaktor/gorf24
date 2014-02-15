package reg

import (
	"testing"

	"github.com/galaktor/gorf24/pipe"
	"github.com/galaktor/gorf24/reg/addr"
)

func TestNewRxPayloadWidth_RegisterAddress_AllPipes_HasRightRegister(t *testing.T) {
	for i := 0; i <= 5; i++ {
		p := pipe.P(i)
		w := NewRxPayloadWidth(p, 0)
		expected := addr.RX_PW(p)
		
		actual := w.Address()

		if actual != expected {
			t.Errorf("expected '%b' but found '%b' with rxpw '%v'", expected, actual, w)
		}
	}
}










