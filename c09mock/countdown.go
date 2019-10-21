package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	countdownStart = 3
	finalWord = "Go!"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main()  {
	//sleeper := &DefaultSleeper{}
	//Countdown(os.Stdout, sleeper)
	configurableSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, configurableSleeper)
}