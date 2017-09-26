package sender

import (
	"github.com/bwmarrin/discordgo"
)

type Sender struct {
    session *discordgo.Session
    channelid string
    user *discordgo.User
}

func NewSender(s *discordgo.Session, m *discordgo.MessageCreate) *Sender{
	this := new(Sender)

	this.session = s
	this.channelid = m.ChannelID
	this.user = m.Author

	return this
}

// func (this Sender) Menu(title string, desc string, options []MenuOption) error { ... }
// func (this Sender) Whisper(msg string) error { ... }