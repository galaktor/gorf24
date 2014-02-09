package reg

type CrcLength byte

const (
	CRC_DISABLED = iota
	CRC_8BIT
	CRC_16BIT
)