package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dcondomitti/go-pivotaltracker/v5/pivotal"
	"github.com/julienschmidt/httprouter"
)

type StoryHandler struct {
	client pivotal.Client
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (sh StoryHandler) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	s, err := sh.Get(params.ByName("story_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sp := NewSlackPayload(&config, s)

	if r.Method == "POST" {
		if room := r.URL.Query().Get("room"); room != "" {
			err = sp.SendMessage(config.SlackWebhookURL, room)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "room parameter required", http.StatusBadRequest)
		}
	} else {
		if err = json.NewEncoder(w).Encode(&sp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (sh *StoryHandler) Get(s string) (*pivotal.Story, error) {
	storyId, err := strconv.Atoi(s)

	story, _, err := sh.client.Stories.Get(storyId)

	if err != nil {
		log.Print(err)
		return &pivotal.Story{}, err
	}

	return story, nil
}
