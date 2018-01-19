package command

import (
	"fmt"

	"github.com/7thFox/hypothesisbot/sender"
)

// Help opens help menus about the commands
type Help struct {
	title    string
	desc     string
	menutree []sender.MenuOption
}

func (c Help) Name() string {
	return "help"
}

func (c Help) HelpText() string {
	return "This menu, dummy"
}

func (c Help) Execute(s sender.Sender, args string) error {
	return s.Menu(c.title, c.desc, c.menutree)
}

func NewHelp(cmds []Command) *Help {
	c := new(Help)
	c.title = "Help Menu"
	c.desc = "Select a command to see it's help message."
	c.menutree = []sender.MenuOption{}

	for _, cmd := range cmds {
		c.menutree = append(c.menutree, sender.MenuSay{
			cmd.Name(),
			"",
			fmt.Sprintf("__**%s**__\n\n%s", cmd.Name(), cmd.HelpText()),
		})
	}

	return c
}
