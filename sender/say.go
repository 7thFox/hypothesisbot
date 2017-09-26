package sender

func (s Sender) Say(msg string) error {
	// _, err := s.session.ChannelMessageSend(s.channelid, "\u200B F"+msg)
	_, err := s.session.ChannelMessageSend(s.channelid, "I'm running the most recent code")
	return err
	// return nil
}
