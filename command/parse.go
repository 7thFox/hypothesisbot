package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!"

// ParseCommand creates a Command object based on the message sent
func ParseCommand(m *discordgo.MessageCreate) (*Command, error) {
	cmd := new(Command)
	*cmd = NewSay(strings.Replace(m.Content, prefix+"say ", "", 1))
	return cmd, nil
}
