package command

import (
	"hypothesisbot/sender"
)

type Say struct {
    msg string
}

func (this Say) Execute(s sender.Sender) error {
	return s.Say(this.msg)
}

func NewSay(s string) *Say {
	this := new(Say)
	this.msg = s
	return this
}