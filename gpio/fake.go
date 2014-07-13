package gpio

type Fake struct {

}

// implement Pin interface
func (f *Fake) Close() error {
	return nil
}

func (f *Fake) Nr() byte {
	return 0
}

func (f *Fake) Direction() Direction {
	return IN
}

func (f *Fake) Set(s PinState) error {
	return nil
}

func (f *Fake) Get() (PinState, error) {
	return LOW, nil
}
