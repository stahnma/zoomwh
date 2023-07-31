package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"time"
)

func sendIRC(message string) {
	ircEnable, _ := os.LookupEnv("ZOOMWH_IRC_ENABLE")
	var (
		ircServer   string
		ircChannel  string
		ircNick     string
		ircPassword string
	)
	if ircEnable == "true" {
		// validate all IRC variables
		validateEnvVars("ZOOMWH_IRC_SERVER")
		validateEnvVars("ZOOMWH_IRC_CHANNEL")
		validateEnvVars("ZOOMWH_IRC_NICK")
		validateEnvVars("ZOOMWH_IRC_PASS")

		ircServer, _ = os.LookupEnv("ZOOMWH_IRC_SERVER")
		ircChannel, _ = os.LookupEnv("ZOOMWH_IRC_CHANNEL")
		ircNick, _ = os.LookupEnv("ZOOMWH_IRC_NICK")
		ircPassword, _ = os.LookupEnv("ZOOMWH_IRC_PASS")
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

	irccon.Privmsg(ircChannel, message)

	time.AfterFunc(1*time.Second, func() {
		irccon.Quit()
	})

	irccon.Loop()
}
