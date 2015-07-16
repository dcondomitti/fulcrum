package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dcondomitti/go-pivotaltracker/v5/pivotal"
)

type SlackPayload struct {
	Attachments []slackMessage `json:"attachments"`
	Channel     string         `json:"channel"`
	Username    string         `json:"username"`
	IconEmoji   string         `json:"icon_emoji"`
}

type slackMessage struct {
	Fallback  string              `json:"fallback"`
	Title     string              `json:"title"`
	TitleLink string              `json:"title_link"`
	Text      string              `json:"text"`
	Fields    []slackMessageField `json:"fields"`
}

type slackMessageField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func NewSlackPayload(c *Config, s *pivotal.Story) SlackPayload {
	return SlackPayload{
		Attachments: extractAttachments(s),
		Username:    c.Username,
		IconEmoji:   c.IconEmoji,
	}
}

func extractAttachments(s *pivotal.Story) []slackMessage {
	return []slackMessage{
		slackMessage{
			Fallback:  s.Name,
			Title:     s.Name,
			TitleLink: s.URL,
			Fields:    extractFields(s),
			Text:      s.Description,
		},
	}
}

func extractFields(s *pivotal.Story) []slackMessageField {

	return []slackMessageField{
		slackMessageField{Title: "State", Value: s.State, Short: true},
		slackMessageField{Title: "Estimate", Value: extractEstimate(s), Short: true},
	}
}

func (sp *SlackPayload) SendMessage(u string, r string) error {
	sp.Channel = r

	encoded, err := sp.Encode()
	if err != nil {
		return err
	}

	resp, err := http.PostForm(u, url.Values{"payload": {encoded}})
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("Not OK")
	}

	return nil
}

func (sp SlackPayload) Encode() (string, error) {
	b, err := json.Marshal(sp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func extractEstimate(s *pivotal.Story) string {
	e := s.Estimate
	if e != nil {
		return fmt.Sprintf("%s points", strconv.FormatFloat(*s.Estimate, 'g', 1, 64))
	}

	return "n/a"
}
