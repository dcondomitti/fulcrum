package main

import (
	"log"
	"net/http"

	"github.com/dcondomitti/fulcrum/middleware"

	"github.com/dcondomitti/go-pivotaltracker/v5/pivotal"
	"github.com/julienschmidt/httprouter"
)

var config Config

func main() {
	log.Print("Starting fulcrum")
	config = parseFlags()

	if config.PivotalAPIKey == "" {
		log.Fatal("Pivotal Tracker API key not present")
	}

	pivotalClient := pivotal.NewClient(config.PivotalAPIKey)

	s := StoryHandler{
		client: *pivotalClient,
	}

	router := httprouter.New()
	router.GET("/projects/:project_id/stories/:story_id", s.Show)
	router.GET("/n/projects/:project_id/stories/:story_id", s.Show)
	router.GET("/stories/:story_id", s.Show)
	router.GET("/story/show/:story_id", s.Show)

	router.POST("/projects/:project_id/stories/:story_id", s.Show)
	router.POST("/n/projects/:project_id/stories/:story_id", s.Show)
	router.POST("/stories/:story_id", s.Show)
	router.POST("/story/show/:story_id", s.Show)

	httpTarget := middleware.Authentication(config.AuthenticationToken, middleware.JsonContentType(router))

	if config.UseTLS {
		listenTLS(config, httpTarget)
	} else {
		listen(config, httpTarget)
	}

}

func listenTLS(c Config, h http.Handler) {
	if c.TLSCertificatePath == "" || c.TLSKeyPath == "" {
		log.Fatal("TLS_CERTIFICATE_PATH and TLS_KEY_PATH are required if USE_TLS is enabled")
	}
	log.Printf("Using certificate %s (%s)", c.TLSCertificatePath, c.TLSKeyPath)
	log.Printf("Starting TLS listener on %s", c.ListenString())
	err := http.ListenAndServeTLS(c.ListenString(), c.TLSCertificatePath, c.TLSKeyPath, h)
	if err != nil {
		panic(err)
	}
}

func listen(c Config, h http.Handler) {
	log.Printf("Starting listener on %s", c.ListenString())
	err := http.ListenAndServe(c.ListenString(), h)
	if err != nil {
		panic(err)
	}
}
