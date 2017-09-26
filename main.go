package main

import (
	"bytes"
	"fmt"
	"hypothesisbot/command"
	"hypothesisbot/sender"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func getToken() string {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("token") // Error handling elided for brevity.
	io.Copy(buf, f)          // Error handling elided for brevity.
	f.Close()
	s := string(buf.Bytes())

	return s
}

func main() {
	discord, err := discordgo.New("Bot " + getToken()) // No more pushing code with my token
	if err != nil {
		fmt.Println("Err ", err)
		return
	}

	fmt.Println("connected")
	discord.AddHandler(messageHandler)

	if err = discord.Open(); err != nil {
		fmt.Println("Err ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	sender := sender.NewSender(s, m)

	if strings.HasPrefix(m.Content, "!say") {
		cmd, _ := command.ParseCommand(m)
		(*cmd).Execute(*sender)
	}
}
