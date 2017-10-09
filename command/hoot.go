package command

import (
	"time"

	"github.com/7thFox/hypothesisbot/sender"
	"github.com/bwmarrin/discordgo"
)

type Hoot struct {
}

func (c Hoot) Execute(s sender.Sender, d *discordgo.Session) error {
	msg := make(chan string, 10)
	exit := make(chan int)
	go s.Update(msg, exit)
	go func() {
		msg <- "pff owls are lame"
		time.Sleep(1000 * time.Millisecond)
		msg <- "Hoot Hoot!"
		exit <- 0
	}()
	return nil
}

func NewHoot(args string) *Hoot {
	this := new(Hoot)
	return this
}
