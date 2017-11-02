package sqlite

import "github.com/bwmarrin/discordgo"
import "fmt"

type Sqlite struct {
}

func NewSqlite(l string) *Sqlite {
	db := Sqlite{}
	return &db
}

func (db *Sqlite) LogMessage(m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
}
