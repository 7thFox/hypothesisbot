package startup

import (
	"fmt"
	"strings"
	"time"

	"github.com/7thFox/hypothesisbot/src/database"
	"github.com/7thFox/hypothesisbot/src/log"

	"github.com/bwmarrin/discordgo"
)

func LogServerFast(svrs []string, startTime time.Time, d *discordgo.Session, db database.Database, lgr log.Logger) error {
	var err error
	lgr.LogState("Logging new messages")
	newMsgs, _ := db.NewestMessagesBefore(startTime)

	for _, s := range svrs {
		chans, _ := d.GuildChannels(s)
		for _, ch := range chans {
			lgr.LogState(fmt.Sprintf("Checking %s #%s", s, ch.Name))
			if newMsgs[ch.ID] < ch.LastMessageID {
				lgr.LogState(fmt.Sprintf("Logging  %s #%s", s, ch.Name))
				lastMsg := ""
				for msgs, err := d.ChannelMessages(ch.ID, 100, "", "", ""); err == nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch.ID, 100, lastMsg, "", "") {
					for _, m := range msgs {
						lastMsg = m.ID
						if strings.Compare(m.ID, newMsgs[ch.ID]) < 0 {
							break
						}
						if !db.IsLogged(m.ID) {
							db.LogMessage(m)
						}
					}
					if strings.Compare(lastMsg, newMsgs[ch.ID]) < 0 {
						break
					}
				}
			}
		}
	}
	lgr.LogState("New messages logged")
	return err
}
