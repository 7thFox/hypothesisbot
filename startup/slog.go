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
		if err := logChannelFull(ch.ID, d, db); err != nil {
			return err
		}
	}
	return nil
}

func logChannelFull(ch string, d *discordgo.Session, db database.Database) error {
	err := logChannelOld(ch, d, db)
	if err != nil {
		return err
	}
	err = logChannelNew(ch, d, db)
	return err
}

func logChannelOld(ch string, d *discordgo.Session, db database.Database) error {
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

func logChannelNew(ch string, d *discordgo.Session, db database.Database) error {
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
