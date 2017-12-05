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

/************  Startup Flags  ************/
var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var slog = flag.String("slog", "", "log all channels of given server")
var configPath = flag.String("config", "./config.json", "set location of config file")

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

	cfg = config.NewConfig(*configPath, *debugMode)

	discord, err := discordgo.New("Bot " + cfg.Token()) // No more pushing code with my token
	if err != nil {
		fmt.Println("Err ", err)
		return
	}

	fmt.Println("connected")

	if *slog != "" {
		fmt.Println("Server Log Mode enabled: Logging...")
		//138977883036188672
		logServer(*slog, discord)
		fmt.Println("Finished Logging Server")
		discord.Close()
		os.Exit(0)
	}

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

func logServer(s string, d *discordgo.Session) {
	chans, _ := d.GuildChannels(s)
	for _, ch := range chans {
		fmt.Printf("Logging %s: %s\n", ch.ID, ch.Name)
		logChannel(ch.ID, d)
	}
}

func logChannel(ch string, d *discordgo.Session) {
	lastMsg := "" //"346148288803897355"
	var err error
	for msgs, err := d.ChannelMessages(ch, 100, lastMsg, "", ""); err == nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch, 100, lastMsg, "", "") {
		for _, m := range msgs {
			lastMsg = m.ID
			if !cfg.Database().IsLogged(m.ID) {
				fmt.Printf("\rLogging ID %s", m.ID)
				cfg.Database().LogMessage(m)
			}
		}
	}
	if err != nil {
		fmt.Println("logChannel", err.Error())
	}

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sender := sender.NewSender(s, m)
	cfg.Database().LogMessage(m.Message)

	cmd, _ := command.ParseCommand(m, prefix(), *debugMode)
	if cmd != nil {
		(*cmd).Execute(*sender, s)
	}
}
