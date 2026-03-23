package mocking

import (
	"fmt"
	"io"
	"time"
)

const(
	finalWord = "Go!"
	countdownStart = 3
)

type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func NewConfigurableSleeper(d time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
    return &ConfigurableSleeper{d, sleep}
}
type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, s Sleeper){
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out,"%d\n",i)
		s.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}
