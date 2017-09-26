package sender

// Say simply outputs the text with a 0-width character at the beginning
func (s Sender) Say(msg string) error {
	_, err := s.session.ChannelMessageSend(s.channelid, "\u2063"+msg)
	return err
}
