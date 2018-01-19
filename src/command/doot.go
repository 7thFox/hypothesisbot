package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/src/sender"
)

type Doot struct {
}

func (c Doot) Name() string {
	return "doot"
}
func (c Doot) HelpText() string {
	return "Provides calcium"
}

func (c Doot) Execute(s sender.Sender, args string) error {
	msg := make(chan string, 10)
	exit := make(chan int)
	go s.Update(msg, exit)
	go func() {
		msg <- "<:doot:254265194732191744> :door:"
		time.Sleep(2500 * time.Millisecond)
		msg <- "<:doot:254265194732191744> <:doot:254265194732191744>"
		exit <- 0
	}()
	return nil
}

func NewDoot() *Doot {
	this := new(Doot)
	return this
}
