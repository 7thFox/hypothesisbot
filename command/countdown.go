package command

import (
	"fmt"
	"time"

	"github.com/7thFox/hypothesisbot/sender"
)

// Countdown updates the message counting down from a number
type Countdown struct {
}

func (c Countdown) Name() string {
	return "countdown"
}

func (c Countdown) HelpText() string {
	return "Countdown updates the message counting down from 10"
}

func (c Countdown) Execute(s sender.Sender, args string) error {
	msg := make(chan string, 10)
	exit := make(chan int)
	go s.Update(msg, exit)
	go func() {
		for i := 5; i > 0; i-- {
			msg <- fmt.Sprintf("%d", i)
			time.Sleep(1000 * time.Millisecond)
		}
		msg <- "Countdown Finished."
		exit <- 0
	}()
	return nil
}

func NewCountdown() *Countdown {
	c := new(Countdown)
	return c
}
