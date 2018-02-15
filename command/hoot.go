package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/sender"
)

type Hoot struct {
}

func (c Hoot) Name() string {
	return "hoot"
}
func (c Hoot) HelpText() string {
	return "For owls and owl enthusiasts alike"
}

func (c Hoot) Execute(s sender.Sender, args string) error {
	msg := make(chan string, 10)
	exit := make(chan int)
	go s.Update(msg, exit)
	go func() {
		msg <- "pff owls are lame"
		time.Sleep(1000 * time.Millisecond)
		msg <- "Hoot Hoot!"
		exit <- 0
	}()
	return nil
}

func NewHoot() *Hoot {
	this := new(Hoot)
	return this
}
