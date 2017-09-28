package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/sender"
)

type Doot struct {
}

func (c Doot) Execute(s sender.Sender) error {
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

func NewDoot(args string) *Doot {
	this := new(Doot)
	return this
}
