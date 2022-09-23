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

const finalWord = "Go!"
const countdownStart = 3

func main() {
	dSleeper := DefaultSleeper{}
	Countdown(os.Stdout, dSleeper)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprintf(w, finalWord)
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
