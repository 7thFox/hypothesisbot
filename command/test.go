package command

import (
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

// Test sends a test message
type Test struct {
}

func (t Test) Execute(s sender.Sender, d *discordgo.Session) error {
	return s.Say("Hello World!")
}

func NewTest(args string) *Test {
	this := new(Test)
	return this
}
