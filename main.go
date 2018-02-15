package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/config"
	"github.com/7thFox/hypothesisbot/database"
	"github.com/7thFox/hypothesisbot/database/mongo"
	"github.com/7thFox/hypothesisbot/log"
	"github.com/7thFox/hypothesisbot/sender"
	"github.com/7thFox/hypothesisbot/startup"
	"github.com/7thFox/hypothesisbot/web"

	"github.com/bwmarrin/discordgo"
)

/************  Startup Flags  ************/
var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")
var configPath = flag.String("config", "./config.json", "set location of config file")

var cfg *config.Config
var logr log.Logger
var db database.Database
var cmds map[string]command.Command

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Loading config from %s\n", *configPath)
	cfg, err = config.NewConfig(*configPath, *debugMode)
	handleError(err)
	fmt.Println("Config loaded")

	db, err = makeDatabase()
	handleError(err)

	fmt.Println("Creating session")
	session, err := discordgo.New("Bot " + cfg.Token())
	handleError(err)
	handleError(session.Open())
	defer session.Close()

	fmt.Println("Creating logger")
	logr = makeLogger(session)

	logr.Log("Connected")
	if *debugMode {
		logr.Log("Debug mode enabled")
	}

	go web.StartWeb(db)

	cmds = commands()
	session.AddHandler(messageHandler)
	startup.LogServerFast(cfg.LogServers(), cfg.StartTime, session, db, logr)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
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

func makeDatabase() (database.Database, error) {
	fmt.Printf("Connecting to %s db at %s.%s\n", cfg.DatabaseType(), cfg.DatabaseHost(), cfg.DatabaseName())
	switch cfg.DatabaseType() {
	case "mongo":
		fmt.Println("Mongo DB connected")
		return mongo.NewMongo(cfg.DatabaseHost(), cfg.DatabaseName())
	}
	return nil, errors.New("Unsupported Database type")
}

func makeLogger(session *discordgo.Session) log.Logger {
	lgr := log.NewMultiLogger()

	if cfg.LogConsole() {
		lgr.Attach(log.NewConsoleLogger())
	}

	if n := cfg.LogDbName(); n != "" {
		// TODO
	}

	if cfg.LogChannelID() != "" {
		lgr.Attach(log.NewChannelLogger(session, cfg.LogChannelID()))
	}

	return lgr
}

func commands() map[string]command.Command {
	cmdsAll := []command.Command{
		command.NewDoot(),
		command.NewGit(),
		command.NewHoot(),
		command.NewKill(),
		command.NewNoot(),
		command.NewSay(),
		command.NewTest(),
		command.NewVersion(cfg.Version),
	}

	cmdsAll = append(cmdsAll, command.NewHelp(cmdsAll))

	// read config and construct commands
	cmds := map[string]command.Command{}
	for _, cmd := range cmdsAll {
		if !cfg.CommandBlacklist()[cmd.Name()] {
			cmds[cmd.Name()] = cmd
		}
	}

	return cmds
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	db.LogMessage(m.Message)
	if !isCmd(m.Content) {
		return
	}

	str := strings.SplitN(m.Content, " ", 2)
	str[0] = strings.TrimPrefix(str[0], cfg.Prefix())

	if cmd := cmds[str[0]]; cmd != nil {
		sender := sender.NewSender(s, m, logr)
		if len(str) > 1 {
			cmd.Execute(*sender, str[1])
		} else {
			cmd.Execute(*sender, "")
		}
	}
}
