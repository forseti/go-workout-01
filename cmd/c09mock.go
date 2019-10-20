package main

import (
	"forseti.github.io/goworkout/c09mock"
	"os"
)

func main()  {
	sleeper := &c09mock.DefaultSleeper{}
	c09mock.Countdown(os.Stdout, sleeper)
}
