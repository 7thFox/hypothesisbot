package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/config"
	"github.com/7thFox/hypothesisbot/sender"

	"github.com/bwmarrin/discordgo"
)

var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var configPath = "./config.json"
var cfg *config.Config

func prefix() string {
	if *debugMode {
		return "!!"
	}
	return "!"
}

func main() {
	flag.Parse()
	if *debugMode {
		fmt.Println("Debug Mode")
	}

	cfg = config.NewConfig(configPath, *debugMode)

	discord, err := discordgo.New("Bot " + cfg.Token()) // No more pushing code with my token
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
	sender := sender.NewSender(s, m)
	cfg.Database().LogMessage(m)

	cmd, _ := command.ParseCommand(m, prefix(), *debugMode)
	if cmd != nil {
		(*cmd).Execute(*sender, s)
	}
}
