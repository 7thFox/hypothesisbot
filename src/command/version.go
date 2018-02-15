package command

import (
	"fmt"

	"github.com/7thFox/hypothesisbot/src/sender"
)

// Say repeats text back into chat
type Version struct {
	version string
}

func (c Version) Name() string {
	return "version"
}
func (c Version) HelpText() string {
	return "Outputs the version of the bot currently running"
}

func (c Version) Execute(s sender.Sender, args string) error {
	return s.Say(fmt.Sprintf("Currently running version %s", c.version))
}

func NewVersion(v string) *Version {
	c := new(Version)
	c.version = v
	return c
}
