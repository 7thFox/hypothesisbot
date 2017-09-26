package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const PREFIX = "!"

func ParseCommand(m *discordgo.MessageCreate) (*Command, error) {
	cmd := new(Command)
	*cmd = NewSay(strings.Replace(m.Content, PREFIX+"say", "", 1))
	return cmd, nil
}