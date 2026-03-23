package counter

import "errors"

type Counter struct{
	value int
}
var ErrCounterAtZero = errors.New("Cannot Decrement Below Zero")

func (c *Counter) Increment()  {
	c.value ++
}
func (c *Counter) Decrement() error {
	if c.value <= 0{
		return ErrCounterAtZero
	}
	c.value --
	return  nil
}
func (c Counter) Value() int {
	return c.value
}