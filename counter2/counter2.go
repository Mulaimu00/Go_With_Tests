package counter2

import "errors"

type Counter struct {
	value int
}

var ErrcounterAtZero = errors.New("Cant decrease 0 or number below it")

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() error{
	if c.value <= 0 {
		return ErrcounterAtZero
	}
	c.value --
	return nil
}

func (c Counter) Value() int {
	return c.value
}