package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/src/sender"
)

type Noot struct {
}

func (c Noot) Name() string {
	return "noot"
}

func (c Noot) HelpText() string {
	return "For people that understand that penguins are better than owls"
}

func (c Noot) Execute(s sender.Sender, args string) error {
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

func NewNoot() *Noot {
	this := new(Noot)
	return this
}
