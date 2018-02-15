package command

import (
	"github.com/7thFox/hypothesisbot/src/sender"
)

type Command interface {
	Execute(s sender.Sender, args string) error
	Name() string
	HelpText() string
}
