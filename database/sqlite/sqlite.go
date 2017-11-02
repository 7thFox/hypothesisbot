package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func NewSqlite(l string) *Sqlite {
	ls := Sqlite{}
	db, err := sql.Open("sqlite3", l)
	if err != nil {
		fmt.Println(err.Error())
	}
	ls.db = db

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS foo (s STRING);")
	if err != nil {
		fmt.Println(err)
	}

	return &ls
}

func (db *Sqlite) LogMessage(m *discordgo.MessageCreate) {

	tx, err := db.db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare("INSERT INTO foo (s) VALUES (?);")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	stmt.Exec(m.Content)
	tx.Commit()
}
