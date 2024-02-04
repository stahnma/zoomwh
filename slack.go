package main

import (
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func postToSlackWebHook(msg string) {
	log.Debugln("(postToSlackWebHook) Sending slack message: " + "'" + msg + "'")

	slack_webhook_uri := viper.GetString("slack_webhook_uri")
	log.Debugln("(postToSlackWebHook) slack_webhook_uri:", slack_webhook_uri)
	data := url.Values{
		"payload": {"{\"text\": \"" + msg + "\"}"},
	}
	resp, err := http.PostForm(slack_webhook_uri, data)
	if err != nil {
		log.Errorln("Error posting to slack:", err)
	}
	log.Debugln(resp.Status)
}
