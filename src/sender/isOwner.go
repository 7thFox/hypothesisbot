package sender

// IsOwner returns weather the message was sent by the bot owner
func (s Sender) IsOwner() bool {
	return s.user.ID == "93398876408524800"
}
