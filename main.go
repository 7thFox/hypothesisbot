package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/config"
	"github.com/7thFox/hypothesisbot/log"
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/7thFox/hypothesisbot/startup"

	"github.com/bwmarrin/discordgo"
)

/************  Startup Flags  ************/
var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var slog = flag.String("slog", "", "log all channels of given server")
var configPath = flag.String("config", "./config.json", "set location of config file")

var cfg *config.Config
var lgr log.Logger

func main() {
	flag.Parse()
	cfg = config.NewConfig(*configPath, *debugMode)
	discord, err := discordgo.New("Bot " + cfg.Token())
	handleError(err)

	handleError(discord.Open())
	defer discord.Close()

	lgr = cfg.Logger(discord)
	lgr.Log("Connected")
	cfg.Database() // Initialize the DB now instead of later

	if *debugMode {
		lgr.Log("Debug Mode Enabled")
	}

	startupTasks(discord)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func startupTasks(d *discordgo.Session) {
	if *slog != "" {
		startup.ServerLog(*slog, d, cfg.Database(), lgr)
		d.Close()
		os.Exit(0)
	} else {
		d.AddHandler(messageHandler)
		startup.LogServerFast(cfg.LogServers(), cfg.StartTime, d, cfg.Database(), lgr)
	}
	lgr.Log("Finished startup")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sender := sender.NewSender(s, m, lgr)
	cfg.Database().LogMessage(m.Message)

	cmd, _ := command.ParseCommand(m, cfg.Prefix(), *debugMode)
	if cmd != nil {
		(*cmd).Execute(*sender, s)
	}
}
