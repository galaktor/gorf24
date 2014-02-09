package reg

/*
Note: Addresses 18 to 1B are reserved for test purposes, altering them will make the chip malfunc-
tion.
*/

type Address byte

func (a Address) Byte() byte {
	return byte(a)
}