package sender

// Say simply outputs the text with a 0-width character at the beginning
func (s Sender) Say(msg string) error {
	_, err := s.session.ChannelMessageSend(s.channelid, "\u2063"+msg)
	return err
}

// SayReplace simply replaces the message with msgID text and a 0-width character at the beginning
func (s Sender) SayReplace(msg string, msgID string) error {
	_, err := s.session.ChannelMessageEdit(s.channelid, msgID, "\u2063"+msg)
	return err
}
