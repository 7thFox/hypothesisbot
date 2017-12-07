package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/config"
	"github.com/7thFox/hypothesisbot/log"
	"github.com/7thFox/hypothesisbot/sender"

	"github.com/bwmarrin/discordgo"
)

/************  Startup Flags  ************/
var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var slog = flag.String("slog", "", "log all channels of given server")
var configPath = flag.String("config", "./config.json", "set location of config file")

var cfg *config.Config
var lgr log.Logger

func main() {
	lgr = log.NewConsoleLogger()
	flag.Parse()
	if *debugMode {
		lgr.Log("Debug Mode")
	}

	cfg = config.NewConfig(*configPath, *debugMode)

	discord, err := discordgo.New("Bot " + cfg.Token()) // No more pushing code with my token
	if err != nil {
		lgr.Log(err.Error())
		return
	}

	lgr.Log("Connected")

	if *slog != "" {
		lgr.Log("Server Log Mode enabled: Logging...")
		logServer(*slog, discord)
		lgr.Log("Finished Logging Server")
		discord.Close()
		os.Exit(0)
	} else {
		logServerFast(discord)
	}

	discord.AddHandler(messageHandler)

	if err = discord.Open(); err != nil {
		lgr.Log(err.Error())
		return
	}

	lgr.Log("Finished startup")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func logServer(s string, d *discordgo.Session) {
	chans, _ := d.GuildChannels(s)
	for _, ch := range chans {
		logChannelFull(ch.ID, d)
	}
}

func logChannelNew(ch string, d *discordgo.Session) {
	lastMsg := ""
	new, _ := cfg.Database().NewestMessageInChannel(ch)
	for msgs, err := d.ChannelMessages(ch, 100, "", "", ""); err == nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch, 100, lastMsg, "", "") {
		for _, m := range msgs {
			lastMsg = m.ID
			if strings.Compare(m.ID, new.ID) < 0 {
				break
			}
			if !cfg.Database().IsLogged(m.ID) {
				cfg.Database().LogMessage(m)
			} else {
			}
		}
		if strings.Compare(lastMsg, new.ID) < 0 {
			break
		}
	}
}

func logChannelOld(ch string, d *discordgo.Session) {
	lastMsg := ""
	old, _ := cfg.Database().OldestMessageInChannel(ch)

	for msgs, err := d.ChannelMessages(ch, 100, old.ID, "", ""); err == nil && len(msgs) > 0; msgs, err = d.ChannelMessages(ch, 100, lastMsg, "", "") {
		for _, m := range msgs {
			lastMsg = m.ID
			if !cfg.Database().IsLogged(m.ID) {
				cfg.Database().LogMessage(m)
			} else {
			}
		}
	}
}

func logServerFast(d *discordgo.Session) {
	lgr.LogState("Logging new messages")
	newMsgs, _ := cfg.Database().NewestMessages()

	for _, s := range cfg.LogServers() {
		chans, _ := d.GuildChannels(s)
		for _, ch := range chans {
			if newMsgs[ch.ID] < ch.LastMessageID {
				logChannelNew(ch.ID, d)
			}
		}
	}
	lgr.Log("Finished scanning channels")
}

func logChannelFull(ch string, d *discordgo.Session) {
	logChannelOld(ch, d)
	logChannelNew(ch, d)
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sender := sender.NewSender(s, m)
	cfg.Database().LogMessage(m.Message)

	cmd, _ := command.ParseCommand(m, cfg.Prefix(), *debugMode)
	if cmd != nil {
		(*cmd).Execute(*sender, s)
	}
}
