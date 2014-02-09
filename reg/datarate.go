package reg

type Datarate byte

const (
	RATE_1MBPS Datarate = iota
	RATE_2MBPS
	RATE_250KBPS
)