package command

import (
	"github.com/7thFox/hypothesisbot/src/sender"
)

// Test sends a test message
type Test struct {
}

func (c Test) Name() string {
	return "test"
}
func (c Test) HelpText() string {
	return "Simple function to test if the bot's online"
}

func (t Test) Execute(s sender.Sender, args string) error {
	return s.Say("Hello World!")
}

func NewTest() *Test {
	this := new(Test)
	return this
}
