package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type dWebhook struct {
	Content   string   `json:"content,omitifempty"`
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Embeds    []embeds `json:"embeds,omitifempty"`
}

type embedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type embedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type embeds struct {
	Fields []embedField `json:"fields"`
	Footer embedFooter  `json:"footer"`
}

func sendWebhook(whurl string, dw dWebhook) error {
	if len(dw.Username) > 32 {
		dw.Username = dw.Username[:32]
	}

	data, err := json.Marshal(&dw)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(whurl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode/100 != 2 {
		io.Copy(os.Stderr, resp.Body)
		resp.Body.Close()
		return fmt.Errorf("status code was %v", resp.StatusCode)
	}

	return nil
}
