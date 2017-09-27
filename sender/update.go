package sender

import (
	"github.com/bwmarrin/discordgo"
)

// Update receives strings and updates the message until the exit signal is received
func (s Sender) Update(m chan string, ex chan int) {
	var msg *discordgo.Message
	for {
		select {
		case ms := <-m:
			if msg == nil {
				msg, _ = s.session.ChannelMessageSend(s.channelid, ms)
			} else {
				s.session.ChannelMessageEdit(s.channelid, msg.ID, ms)
			}
		case <-ex:
			return
		}
	}
}
