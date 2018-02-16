package mongo

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type storeMessage struct {
	ID              string                         `json:"id"`
	ChannelID       string                         `json:"channel_id"`
	Content         string                         `json:"content"`
	Timestamp       time.Time                      `json:"timestamp"`
	EditedTimestamp time.Time                      `json:"edited_timestamp"`
	MentionRoles    []string                       `json:"mention_roles"`
	Tts             bool                           `json:"tts"`
	MentionEveryone bool                           `json:"mention_everyone"`
	Author          *discordgo.User                `json:"author"`
	Attachments     []*discordgo.MessageAttachment `json:"attachments"`
	Embeds          []*discordgo.MessageEmbed      `json:"embeds"`
	Mentions        []*discordgo.User              `json:"mentions"`
	Reactions       []*discordgo.MessageReactions  `json:"reactions"`
	Type            discordgo.MessageType          `json:"type"`
}

func (db *Mongo) LogMessage(m *discordgo.Message) error {
	ts, _ := time.Parse(time.RFC3339Nano, string(m.Timestamp))
	ets, _ := time.Parse(time.RFC3339Nano, string(m.EditedTimestamp))
	mm := storeMessage{
		m.ID,
		m.ChannelID,
		m.Content,
		ts,
		ets,
		m.MentionRoles,
		m.Tts,
		m.MentionEveryone,
		m.Author,
		m.Attachments,
		m.Embeds,
		m.Mentions,
		m.Reactions,
		m.Type,
	}
	err := db.messages.Insert(mm)
	return err
}
