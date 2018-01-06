package sender

import (
	"github.com/bwmarrin/discordgo"

	"github.com/7thFox/hypothesisbot/log"
)

// Sender is an object for basic replies in channels
type Sender struct {
	session   *discordgo.Session
	channelid string
	user      *discordgo.User
	msgid     string
	logger    log.Logger
}

func NewSender(s *discordgo.Session, m *discordgo.MessageCreate, l log.Logger) *Sender {
	this := new(Sender)

	this.session = s
	this.channelid = m.ChannelID
	this.user = m.Author
	this.msgid = m.ID
	this.logger = l

	return this
}

// func (this Sender) Menu(title string, desc string, options []MenuOption) error { ... }
// func (this Sender) Whisper(msg string) error { ... }
