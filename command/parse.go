package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!"

// ParseCommand creates a Command object based on the message sent
func ParseCommand(m *discordgo.MessageCreate) (*Command, error) {
	// cmd := new(Command)
	var cmd Command

	fmt.Println("parse")

	if isCommand(m, "say") {
		fmt.Println("say")
		cmd = NewSay(stripCommand(m, "say"))
	} else if isCommand(m, "test") {
		fmt.Println("test")
		cmd = NewTest()
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
