/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

/* used for Rx and Tx addresses; hence 'x' address
   max 5 bytes
   using 64bits for simple storage, but will truncate
   first 3 bytes to limit to 5 bytes

   64bits
   <-- ignored -----------><-- used ------------------------------>
   xxxxxxxxxxxxxxxxxxxxxxxx1111111111111111111111111111111111111111 */
type A [5]byte

func New(flags [5]byte) A {
	return A(flags)
}

func NewFromI(flags uint64) A {
	// iToB5 will shift out first 3 bytes, effectively masking them out
	return New(iToB5(flags))
}

func (a A) ToB5() [5]byte {
	return [5]byte(a)
}

func (a A) ToI() uint64 {
	return b5ToI(a.ToB5())
}

func iToB5(bits uint64) [5]byte {
	return [5]byte{
		byte(bits >> 32),
		byte(bits >> 24),
		byte(bits >> 16),
		byte(bits >> 8),
		byte(bits),
	}
}

func b5ToI(b [5]byte) uint64 {
	return uint64(b[0]) << 32 |
	       uint64(b[1]) << 24 |
	       uint64(b[2]) << 16 |
	       uint64(b[3]) << 8 |
	       uint64(b[4])
}
