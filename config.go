package main

import (
	"fmt"

	"github.com/ianschenck/envflag"
)

const (
	DEFAULT_LISTEN_HOST          = "0.0.0.0"
	DEFAULT_LISTEN_PORT          = 28443
	DEFAULT_TLS_KEY_PATH         = "./certificates/key.pem"
	DEFAULT_TLS_CERTIFICATE_PATH = "./certificates/cert.pem"
)

type Config struct {
	PivotalAPIKey      string
	ListenHost         string
	ListenPort         int
	SlackWebhookURL    string
	TLSCertificatePath string
	TLSKeyPath         string
	UseTLS             bool
}

func (c *Config) ListenString() string {
	return fmt.Sprintf("%s:%v", c.ListenHost, c.ListenPort)
}

var (
	pivotalAPIKey      = envflag.String("PIVOTAL_API_KEY", "", "pivotal tracker api key")
	listenHost         = envflag.String("HOST", DEFAULT_LISTEN_HOST, fmt.Sprintf("interface to listen on, default: %v", DEFAULT_LISTEN_HOST))
	listenPort         = envflag.Int("PORT", DEFAULT_LISTEN_PORT, fmt.Sprintf("port to listen on, default: %v", DEFAULT_LISTEN_PORT))
	slackWebhookURL    = envflag.String("SLACK_WEBHOOK_URL", "", "slack incoming webhook url")
	tlsCertificatePath = envflag.String("TLS_CERTIFICATE_PATH", DEFAULT_TLS_CERTIFICATE_PATH, fmt.Sprintf("path to the x509 certificate, default: %v", DEFAULT_TLS_CERTIFICATE_PATH))
	tlsKeyPath         = envflag.String("TLS_KEY_PATH", DEFAULT_TLS_KEY_PATH, fmt.Sprintf("path to x509 private key, default: %v", DEFAULT_TLS_KEY_PATH))
	useTLS             = envflag.Bool("USE_TLS", false, "enable TLS listener, requires TLS_CERTIFICATE_PATH, TLS_PRIVATE_KEY_PATH")
)

func parseFlags() Config {
	envflag.Parse()

	cfg := Config{
		ListenHost:         *listenHost,
		ListenPort:         *listenPort,
		PivotalAPIKey:      *pivotalAPIKey,
		SlackWebhookURL:    *slackWebhookURL,
		TLSCertificatePath: *tlsCertificatePath,
		TLSKeyPath:         *tlsKeyPath,
		UseTLS:             *useTLS,
	}

	return cfg
}
