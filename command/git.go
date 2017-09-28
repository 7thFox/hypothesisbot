package command

import (
	"github.com/7thFox/hypothesisbot/sender"
)

type Git struct {
}

func (t Git) Execute(s sender.Sender) error {
	return s.Say("Source code can be found at https://github.com/7thFox/hypothesisbot")
}

func NewGit(args string) *Git {
	this := new(Git)
	return this
}
