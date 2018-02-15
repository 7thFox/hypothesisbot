package sender

import (
	"fmt"
)

// Reply prepends an @mention at the beginning of the message
func (s Sender) Reply(msg string) error {
	_, err := s.session.ChannelMessageSend(s.channelid, fmt.Sprintf("<@%s> %s", s.user.ID, msg))
	return err
}
