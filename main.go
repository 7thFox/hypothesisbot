package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/7thFox/hypothesisbot/command"
	"github.com/7thFox/hypothesisbot/sender"

	"github.com/bwmarrin/discordgo"
)

var debugMode = flag.Bool("debug", false, "run in debug mode with debug settings")

func tokenFilename() string {
	if *debugMode {
		return "token.debug"
	}
	return "token"
}

func prefix() string {
	if *debugMode {
		return "!!"
	}
	return "!"
}

func getToken() string {

	buf := bytes.NewBuffer(nil)
	f, err := os.Open(tokenFilename())
	if err != nil {
		fmt.Println("Could not open token file")
		os.Exit(1)
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		fmt.Println(err)
		f.Close()
		os.Exit(1)
	}
	f.Close()
	s := string(buf.Bytes())

	return s
}

func main() {
	flag.Parse()
	if *debugMode {
		fmt.Println("Debug Mode")
	}

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
	sender := sender.NewSender(s, m)

	cmd, _ := command.ParseCommand(m, prefix(), *debugMode)
	if cmd != nil {
		(*cmd).Execute(*sender, s)
	}
}
