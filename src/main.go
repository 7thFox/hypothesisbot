package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/7thFox/hypothesisbot/src/config"
	"github.com/7thFox/hypothesisbot/src/log"
	"github.com/7thFox/hypothesisbot/src/sender"
	"github.com/7thFox/hypothesisbot/src/startup"
	"github.com/7thFox/hypothesisbot/src/web"

	"github.com/bwmarrin/discordgo"
)

/************  Startup Flags  ************/
var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var slog = flag.String("slog", "", "log all channels of given server")
var cpurge = flag.String("cpurge", "", "purges channels from given server(s) after given date (yyyy-mm-dd)")
var cpurgelist = flag.String("cpurgelist", "", "lists purge canidate channels from given server(s) after given date (yyyy-mm-dd)")
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

	go web.StartWeb()

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
	}

	if *cpurgelist != "" {
		args := strings.Split(*cpurgelist, " ")
		if len(args) < 2 {
			handleError(fmt.Errorf("cpurgelist: Expected 2+ args. Got %d", len(args)))
		}
		t, err := time.Parse("2006-01-02", args[len(args)-1])
		handleError(err)

		for _, sid := range args[:len(args)-1] {
			handleError(startup.ChannelPurgeList(sid, t, d, lgr))
		}

		lgr.Log("End of purge canidates")
		d.Close()
		os.Exit(0)
	}

	if *cpurge != "" {
		args := strings.Split(*cpurge, " ")
		if len(args) < 2 {
			handleError(fmt.Errorf("cpurge: Expected 2+ args. Got %d", len(args)))
		}
		t, err := time.Parse("2006-01-02", args[len(args)-1])
		handleError(err)

		for _, sid := range args[:len(args)-1] {
			handleError(startup.ChannelPurge(sid, t, d, lgr))
		}

		lgr.Log("End of channel purge")
		d.Close()
		os.Exit(0)
	}

	d.AddHandler(messageHandler)
	startup.LogServerFast(cfg.LogServers(), cfg.StartTime, d, cfg.Database(), lgr)
	lgr.Log("Finished startup")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func isCmd(m string) bool {
	return strings.HasPrefix(m, cfg.Prefix())
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg.Database().LogMessage(m.Message)
	if !isCmd(m.Content) {
		return
	}

	str := strings.SplitN(m.Content, " ", 2)
	str[0] = strings.TrimPrefix(str[0], cfg.Prefix())

	if cmd := cfg.Commands()[str[0]]; cmd != nil {
		sender := sender.NewSender(s, m, lgr)
		if len(str) > 1 {
			cmd.Execute(*sender, str[1])
		} else {
			cmd.Execute(*sender, "")
		}
	}
}
