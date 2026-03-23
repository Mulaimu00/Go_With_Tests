package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const(
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls [] string
}
type SpyTime struct {
	durationSlept time.Duration
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}
func TestCountdown(t *testing.T) {
	t.Run("Countdown at 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want :="3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before evry print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}