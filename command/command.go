package command

import (
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Execute(s sender.Sender, d *discordgo.Session) error
}
