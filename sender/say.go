package sender

// import (
// 	"github.com/bwmarrin/discordgo"
// )

func (this Sender) Say(msg string) error {
	_, err := this.session.ChannelMessageSend(this.channelid, msg)
	return err
}
