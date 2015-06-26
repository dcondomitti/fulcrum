package main

import (
	"fmt"

	"github.com/ianschenck/envflag"
)

const (
	DEFAULT_LISTEN_HOST = "0.0.0.0"
	DEFAULT_LISTEN_PORT = 28443
)

type Config struct {
	PivotalAPIKey      string
	ListenHost         string
	ListenPort         int
	TLSCertificatePath string
	TLSKeyPath         string
}

func (c *Config) ListenString() string {
	return fmt.Sprintf("%s:%v", c.ListenHost, c.ListenPort)
}

var (
	pivotalAPIKey = envflag.String("PIVOTAL_API_KEY", "", "pivotal tracker api key")
	listenHost    = envflag.String("HOST", DEFAULT_LISTEN_HOST, fmt.Sprintf("interface to listen on, default: %v", DEFAULT_LISTEN_HOST))
	listenPort    = envflag.Int("PORT", DEFAULT_LISTEN_PORT, fmt.Sprintf("port to listen on, default: %v", DEFAULT_LISTEN_PORT))
)

func parseFlags() Config {
	envflag.Parse()

	cfg := Config{
		ListenHost:    *listenHost,
		ListenPort:    *listenPort,
		PivotalAPIKey: *pivotalAPIKey,
	}

	return cfg
}
