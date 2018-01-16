package startup

import (
	"fmt"
	"time"

	"github.com/7thFox/hypothesisbot/log"
	"github.com/bwmarrin/discordgo"
)

func ChannelPurgeList(sid string, t time.Time, s *discordgo.Session, l log.Logger) error {
	chans, err := s.GuildChannels(sid)
	if err != nil {
		fmt.Println("error getting channels")
		return err
	}
	for _, ch := range chans {
		if ch.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		ms, err := s.ChannelMessages(ch.ID, 1, "", "", "")
		if err != nil || len(ms) < 1 {
			fmt.Println("error getting messages")
			return fmt.Errorf("%s; ms len: %d", err.Error(), len(ms))
		}

		if tt, _ := ms[0].Timestamp.Parse(); tt.Before(t) {
			l.Log(fmt.Sprintf("Purge Canidate: (%s) #%s", ch.ID, ch.Name))
		}
	}

	return nil
}
