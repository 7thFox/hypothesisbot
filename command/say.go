package command

import (
	"github.com/7thFox/hypothesisbot/sender"
)

// Say repeats text back into chat
type Say struct {
	msg string
}

func (sy Say) Execute(s sender.Sender) error {
	return s.Say(sy.msg)
}

func NewSay(s string) *Say {
	this := new(Say)
	this.msg = s
	return this
}
