package command

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!"

// ParseCommand creates a Command object based on the message sent
func ParseCommand(m *discordgo.MessageCreate) (*Command, error) {
	var cmd Command

	if isCommand(m, "say") {
		cmd = NewSay(stripCommand(m, "say"))
	} else if isCommand(m, "test") {
		cmd = NewTest()
	} else if isCommand(m, "countdown") {
		// cmd = NewCountdown()
	}

	if cmd != nil {
		return &cmd, nil
	}
	return nil, errors.New("Not a command message")
}

func isCommand(m *discordgo.MessageCreate, c string) bool {
	return strings.HasPrefix(m.Content, prefix+c)
}

func stripCommand(m *discordgo.MessageCreate, c string) string {
	return strings.TrimSpace(strings.TrimPrefix(m.Content, prefix+c))
}
