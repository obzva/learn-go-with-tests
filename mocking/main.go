package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	count     = 3
	finalWord = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

// func main() {
// 	s := &ConfigurableSleeper{time.Second, time.Sleep}
// 	Countdown(os.Stdout, s)
// }
