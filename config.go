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
	AuthenticationToken string
	PivotalAPIKey       string
	ListenHost          string
	ListenPort          int
	IconEmoji           string
	SlackWebhookURL     string
	TLSCertificatePath  string
	TLSKeyPath          string
	Username            string
	UseTLS              bool
}

func (c *Config) ListenString() string {
	return fmt.Sprintf("%s:%v", c.ListenHost, c.ListenPort)
}

var (
	authenticationToken = envflag.String("AUTHENTICATION_TOKEN", "", "authentication token")
	pivotalAPIKey       = envflag.String("PIVOTAL_API_KEY", "", "pivotal tracker api key")
	listenHost          = envflag.String("HOST", DEFAULT_LISTEN_HOST, fmt.Sprintf("interface to listen on, default: %v", DEFAULT_LISTEN_HOST))
	listenPort          = envflag.Int("PORT", DEFAULT_LISTEN_PORT, fmt.Sprintf("port to listen on, default: %v", DEFAULT_LISTEN_PORT))
	iconEmoji           = envflag.String("ICON_EMOJI", "", "emoji to use for message icon")
	slackWebhookURL     = envflag.String("SLACK_WEBHOOK_URL", "", "slack incoming webhook url")
	tlsCertificatePath  = envflag.String("TLS_CERTIFICATE_PATH", DEFAULT_TLS_CERTIFICATE_PATH, fmt.Sprintf("path to the x509 certificate, default: %v", DEFAULT_TLS_CERTIFICATE_PATH))
	tlsKeyPath          = envflag.String("TLS_KEY_PATH", DEFAULT_TLS_KEY_PATH, fmt.Sprintf("path to x509 private key, default: %v", DEFAULT_TLS_KEY_PATH))
	username            = envflag.String("USERNAME", "Pivotal Tracker", "username to send notifications from")
	useTLS              = envflag.Bool("USE_TLS", false, "enable TLS listener, requires TLS_CERTIFICATE_PATH, TLS_PRIVATE_KEY_PATH")
)

func parseFlags() Config {
	envflag.Parse()

	cfg := Config{
		AuthenticationToken: *authenticationToken,
		IconEmoji:           *iconEmoji,
		ListenHost:          *listenHost,
		ListenPort:          *listenPort,
		PivotalAPIKey:       *pivotalAPIKey,
		SlackWebhookURL:     *slackWebhookURL,
		TLSCertificatePath:  *tlsCertificatePath,
		TLSKeyPath:          *tlsKeyPath,
		Username:            *username,
		UseTLS:              *useTLS,
	}

	return cfg
}
