package reg

type Register struct {
	a Address
	flags byte
}

func (c *Register) Byte() byte {
	return c.flags
}


