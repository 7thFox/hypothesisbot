package command

import (
	"hypothesisbot/sender"
)

type Command interface {
	Execute(s sender.Sender) error
}