package sender

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MenuOption interface {
	Execute(s Sender) error
	Title() string
}

var emojis = []string{
	"1⃣",
	"2⃣",
	"3⃣",
	"4⃣",
	"5⃣",
	"6⃣",
	"7⃣",
	"8⃣",
	"9⃣",
	"🔟",
}

func (s Sender) Menu(title string, desc string, options []MenuOption) error {
	if len(options) > len(emojis) {
		return errors.New("Menu has more than supported number of options")
	}
	var buff bytes.Buffer
	optMp := map[string]MenuOption{}
	buff.WriteString(fmt.Sprintf("\n__**%s**__\n_%s_\n\n", title, desc))
	for i, opt := range options {
		optMp[emojis[i]] = opt
		buff.WriteString(fmt.Sprintf("%s - %s\n", emojis[i], opt.Title()))
	}
	msg, err := s.session.ChannelMessageSend(s.channelid, buff.String())
	if err != nil {
		return err
	}

	for i := range options {
		if err := s.session.MessageReactionAdd(s.channelid, msg.ID, emojis[i]); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	var handler interface{}
	handler = func(ses *discordgo.Session, rc *discordgo.MessageReactionAdd) {
		if rc.MessageID == msg.ID && rc.UserID == s.user.ID {
			optMp[rc.Emoji.Name].Execute(s)
			// s.Say(fmt.Sprintf("%s clicked %s", s.user.Username, rc.Emoji.Name))
		} else {
			s.session.AddHandlerOnce(handler)
		}
	}

	s.session.AddHandlerOnce(handler)

	return nil
}
