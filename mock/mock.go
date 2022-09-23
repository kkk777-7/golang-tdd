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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{3 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprintf(w, finalWord)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}
