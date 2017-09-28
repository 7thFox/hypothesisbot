package command

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ParseCommand creates a Command object based on the message sent
// debug: weather to include experimental commands
func ParseCommand(m *discordgo.MessageCreate, prefix string, debug bool) (*Command, error) {
	var cmd Command
	var cmdStr string
	argStr := ""

	s := strings.SplitN(m.Content, " ", 2)
	cmdStr = s[0]

	if !strings.HasPrefix(cmdStr, prefix) {
		return nil, errors.New("Not a command message")
	}

	if len(s) > 1 {
		argStr = s[1]
	}

	if cmdStr == prefix+"say" {
		cmd = NewSay(argStr)
	} else if cmdStr == prefix+"test" {
		cmd = NewTest(argStr)
	} else if cmdStr == prefix+"countdown" {
		// cmd = NewCountdown(argStr)
	} else if cmdStr == prefix+"doot" {
		cmd = NewDoot(argStr)
	} else if cmdStr == prefix+"hoot" {
		cmd = NewHoot(argStr)
	} else if cmdStr == prefix+"noot" {
		cmd = NewNoot(argStr)
	} else if cmdStr == prefix+"git" {
		cmd = NewGit(argStr)
	}

	if cmd != nil {
		return &cmd, nil
	}
	return nil, errors.New("Not a command message")
}
