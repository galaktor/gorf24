package reg

type PowerLevel byte

const (
	PA_MIN PowerLevel = iota
	PA_LOW
	PA_HIGH
	PA_MAX
	PA_ERROR // what is this for?
)