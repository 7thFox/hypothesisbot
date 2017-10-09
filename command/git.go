package command

import (
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

type Git struct {
}

func (t Git) Execute(s sender.Sender, d *discordgo.Session) error {
	return s.Say("Source code can be found at https://github.com/7thFox/hypothesisbot")
}

func NewGit(args string) *Git {
	this := new(Git)
	return this
}
