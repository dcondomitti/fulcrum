package main

import (
	"strconv"

	"github.com/dcondomitti/go-pivotaltracker/v5/pivotal"
)

type SlackPayload struct {
	Attachments []slackMessage `json:"attachments"`
}

type slackMessage struct {
	Fallback string              `json:"fallback"`
	Text     string              `json:"text"`
	Fields   []slackMessageField `json:"fields"`
}

type slackMessageField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func NewSlackMessage(s *pivotal.Story) SlackPayload {
	return SlackPayload{Attachments: extractAttachments(s)}
}

func extractAttachments(s *pivotal.Story) []slackMessage {
	return []slackMessage{
		slackMessage{
			Fallback: "",
			Text:     s.Name,
			Fields:   extractFields(s),
		},
	}
}

func extractFields(s *pivotal.Story) []slackMessageField {
	return []slackMessageField{
		slackMessageField{Title: "Current State", Value: s.State},
		slackMessageField{Title: "Estimate", Value: strconv.FormatFloat(*s.Estimate, 'g', 1, 64)},
		slackMessageField{Title: "Kind", Value: s.Kind},
		slackMessageField{Title: "Project", Value: strconv.Itoa(s.ProjectId)},
	}
}
