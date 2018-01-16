package startup

import (
	"strings"

	"github.com/7thFox/hypothesisbot/database"

	"github.com/7thFox/hypothesisbot/log"
	"github.com/bwmarrin/discordgo"
)

func ServerLog(svr string, d *discordgo.Session, db database.Database, lgr log.Logger) error {
	lgr.Log("Server Log Mode enabled: Logging...")
	if err := logServer(svr, d, db, lgr); err != nil {
		return err
	}
	lgr.Log("Finished Logging Server")
	return nil
}

func logServer(s string, d *discordgo.Session, db database.Database, lgr log.Logger) error {
	chans, _ := d.GuildChannels(s)
	for _, ch := range chans {
		lgr.LogState("Logging " + ch.Name)
		if err := logChannelFull(ch.ID, db, d); err != nil {
			return err
		}
	}
	return nil
}

func logChannelFull(ch string, db database.Database, d *discordgo.Session) error {
	err := logChannelOld(ch, db, d)
	if err != nil {
		return err
	}
	err = logChannelNew(ch, db, d)
	return err
}

func logChannelOld(ch string, db database.Database, d *discordgo.Session) error {
	lastMsg := ""
	old, _ := db.OldestMessageInChannel(ch)

	var err error
	for msgs, err := d.ChannelMessages(ch, 100, old.ID, "", ""); err != nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch, 100, lastMsg, "", "") {
		for _, m := range msgs {
			lastMsg = m.ID
			if !db.IsLogged(m.ID) {
				db.LogMessage(m)
			}
		}
	}
	return err
}

func logChannelNew(ch string, db database.Database, d *discordgo.Session) error {
	var err error
	lastMsg := ""
	new, _ := db.NewestMessageInChannel(ch)
	for msgs, err := d.ChannelMessages(ch, 100, "", "", ""); err != nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch, 100, lastMsg, "", "") {
		for _, m := range msgs {
			lastMsg = m.ID
			if strings.Compare(m.ID, new.ID) < 0 {
				break
			}
			if !db.IsLogged(m.ID) {
				db.LogMessage(m)
			}
		}
		if strings.Compare(lastMsg, new.ID) < 0 {
			break
		}
	}
	return err
}
