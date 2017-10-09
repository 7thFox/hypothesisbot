package command

import (
	"os"

	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

type Kill struct {
}

func (t Kill) Execute(s sender.Sender, d *discordgo.Session) error {
	s.Say(":skull:")
	d.Logout()
	os.Exit(0)
	return nil
}

func NewKill(args string) *Kill {
	this := new(Kill)
	return this
}
