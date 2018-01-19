package command

import (
	"github.com/7thFox/hypothesisbot/sender"
)

type Git struct {
}

func (c Git) Name() string {
	return "git"
}

func (c Git) HelpText() string {
	return "Provides the link to where you can see the source of this great project"
}

func (c Git) Execute(s sender.Sender, args string) error {
	return s.Say("Source code can be found at https://github.com/7thFox/hypothesisbot")
}

func NewGit() *Git {
	c := new(Git)
	return c
}
