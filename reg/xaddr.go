package reg

/* used for Rx and Tx addresses; hence 'x' address
   max 5 bytes
   using 64bits for simple storage, but will truncate
   first 3 bytes to limit to 5 bytes

   64bits
   <-- ignored -----------><-- used ------------------------------>
   xxxxxxxxxxxxxxxxxxxxxxxx1111111111111111111111111111111111111111 */
type XAddress uint64

func NewXAddress(flags uint64) XAddress {
	// mask out first 3 bytes
	return XAddress(flags & 0x000000FFFFFFFFFF)
}

func (a XAddress) Byte() uint64 {
	return uint64(a)
}
