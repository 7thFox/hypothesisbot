package log

import (
	"github.com/bwmarrin/discordgo"
)

type ChannelLogger struct {
	session *discordgo.Session
	channel string

	lastMessageID string
}

func NewChannelLogger(session *discordgo.Session, channel string) *ChannelLogger {
	l := ChannelLogger{}
	l.session = session
	l.channel = channel
	l.lastMessageID = ""
	return &l
}

func (l *ChannelLogger) Log(m string) error {
	_, err := l.session.ChannelMessageSend(l.channel, m)

	l.lastMessageID = ""
	return err
}
func (l *ChannelLogger) LogState(m string) error {
	var err error
	if l.lastMessageID == "" {
		var msg *discordgo.Message
		msg, err = l.session.ChannelMessageSend(l.channel, m)

		l.lastMessageID = msg.ID
	} else {
		_, err = l.session.ChannelMessageEdit(l.channel, l.lastMessageID, m)
	}
	return err
}
