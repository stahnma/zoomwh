package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"time"
)

func sendIRC(message string) {

	ircEnable, _ := os.LookupEnv("ZOOM_WH_IRC_ENABLE")
	var (
		ircServer   string
		ircChannel  string
		ircNick     string
		ircPassword string
	)
	if ircEnable == "true" {
		// validate all IRC variables
		validateEnvVars("ZOOM_WH_IRC_SERVER")
		validateEnvVars("ZOOM_WH_IRC_CHANNEL")
		validateEnvVars("ZOOM_WH_IRC_NICK")
		validateEnvVars("ZOOM_WH_IRC_PASS")

		ircServer, _ = os.LookupEnv("ZOOM_WH_IRC_SERVER")
		ircChannel, _ = os.LookupEnv("ZOOM_WH_IRC_CHANNEL")
		ircNick, _ = os.LookupEnv("ZOOM_WH_IRC_NICK")
		ircPassword, _ = os.LookupEnv("ZOOM_WH_IRC_PASS")
	}

	irccon := irc.IRC(ircNick, ircNick)
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.Password = ircPassword
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(ircChannel)
	})

	err := irccon.Connect(ircServer)
	if err != nil {
		log.Fatal("Error connecting to IRC server:", err)
	}
	defer irccon.Quit()

	message = "Hello, IRC world! (this is an automated message)"
	irccon.Privmsg(ircChannel, message)

	time.AfterFunc(1*time.Second, func() {
		irccon.Quit()
	})

	irccon.Loop()
}
