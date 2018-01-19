package sender

// Whisper sends `msg` as a PM to the user
func (s Sender) Whisper(msg string) error {
	ch, err := s.session.UserChannelCreate(s.user.ID)
	if err != nil {
		return err
	}
	_, err = s.session.ChannelMessageSend(ch.ID, msg)
	return err
}
