package sender

// DeleteCommand deletes the message that contained the calling command
func (s Sender) DeleteCommand() error {
	return s.session.ChannelMessageDelete(s.channelid, s.msgid)
}
