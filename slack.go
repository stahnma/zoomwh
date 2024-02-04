package main

import (
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func postToSlackWebHook(msg string) {
	log.Debugln("(postToSlackWebHook) Sending slack message", msg)

	slack_uri := viper.GetString("slack_uri")
	data := url.Values{
		"payload": {"{\"text\": \"" + msg + "\"}"},
	}
	resp, err := http.PostForm(slack_uri, data)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln(resp.Status)
}
