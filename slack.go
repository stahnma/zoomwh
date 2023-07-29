package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func postToSlackWebHook(msg string) {

	validateEnvVars("ZOOMWH_SLACK_WH_URI")

	slack_uri := os.Getenv("ZOOMWH_SLACK_WH_URI")
	data := url.Values{
		"payload": {"{\"text\": \"" + msg + "\"}"},
	}
	resp, err := http.PostForm(slack_uri, data)
	fmt.Println(resp.Status)
	if err != nil {
		log.Fatal(err)
	}
}
