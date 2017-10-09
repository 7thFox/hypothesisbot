package command

import (
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

// Say repeats text back into chat
type Say struct {
	msg string
}

func (sy Say) Execute(s sender.Sender, d *discordgo.Session) error {
	s.DeleteCommand()
	return s.Say(sy.msg)
}

func NewSay(args string) *Say {
	this := new(Say)
	this.msg = args
	return this
}
