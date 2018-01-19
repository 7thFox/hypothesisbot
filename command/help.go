package command

import (
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

func NewHelp() *Help {
	c := new(Help)
	c.title = "Help Menu"
	c.desc = "Currently just a test of the menu feature. Sorry, no help for you."
	c.menutree = []sender.MenuOption{
		sender.Submenu{
			"Foo",
			"Does foo things",
			[]sender.MenuOption{
				sender.MenuSay{
					"Bar",
					"Does bar things",
					"BAR!!!!!",
				},
				sender.MenuSay{
					"Foobar",
					"Does foobar things",
					"FOOBAR!!!!!",
				},
			},
		},
		sender.MenuSay{
			"Baz",
			"Does baz things",
			"BAZ!!!!!",
		},
		sender.MenuSay{
			"Foobaz",
			"Does foobaz things",
			"FOOBAZ!!!!!",
		},
	}
	return c
}
