package main

import (
	"log"
	"net/http"

	"github.com/dcondomitti/fulcrum/middleware"

	"github.com/dcondomitti/go-pivotaltracker/v5/pivotal"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Print("Starting fulcrum")
	config := parseFlags()

	if config.PivotalAPIKey == "" {
		log.Fatal("Pivotal Tracker API key not present")
	}

	pivotalClient := pivotal.NewClient(config.PivotalAPIKey)

	s := StoryHandler{
		client: *pivotalClient,
	}

	router := httprouter.New()
	router.GET("/projects/:project_id/stories/:story_id", s.Show)
	router.GET("/stories/:story_id", s.Show)
	router.GET("/story/show/:story_id", s.Show)

	log.Printf("Starting TLS listener on %s", config.ListenString())
	log.Printf("Using certificate %s (%s)", config.TLSCertificatePath, config.TLSKeyPath)

	httpTarget := middleware.JsonContentType(router)

	err := http.ListenAndServeTLS(config.ListenString(), config.TLSCertificatePath, config.TLSKeyPath, httpTarget)
	if err != nil {
		panic(err)
	}
}
