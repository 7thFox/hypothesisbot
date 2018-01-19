package command

import (
	"github.com/7thFox/hypothesisbot/sender"
)

// Say repeats text back into chat
type Say struct {
}

func (c Say) Name() string {
	return "say"
}
func (c Say) HelpText() string {
	return "Repeats back text"
}

func (sy Say) Execute(s sender.Sender, args string) error {
	s.DeleteCommand()
	return s.Say(args)
}

func NewSay() *Say {
	this := new(Say)
	return this
}
