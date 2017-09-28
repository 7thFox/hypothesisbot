package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/sender"
)

type Noot struct {
}

func (c Noot) Execute(s sender.Sender) error {
	msg := make(chan string, 10)
	exit := make(chan int)
	go s.Update(msg, exit)
	go func() {
		msg <- "penguins shall rule the earth!"
		time.Sleep(1000 * time.Millisecond)
		msg <- "noot noot"
		exit <- 0
	}()
	return nil
}

func NewNoot(args string) *Noot {
	this := new(Noot)
	return this
}
