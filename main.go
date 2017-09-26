package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"hypothesisbot/command"
	"hypothesisbot/sender"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot Mjc1OTgzOTg3NjY1NTM0OTc2.DEbucA.8LUP6jTUXG-wNG206aiNtJuME2o")
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

	sender := sender.NewSender(s,m)

	if strings.HasPrefix(m.Content, "!say") {
		cmd, _ := command.ParseCommand(m)
		(*cmd).Execute(*sender)
	}
}