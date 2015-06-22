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
	TLSCertificatePath string
	TLSKeyPath         string
}

func (c *Config) ListenString() string {
	return fmt.Sprintf("%s:%v", c.ListenHost, c.ListenPort)
}

var (
	pivotalAPIKey      = envflag.String("PIVOTAL_API_KEY", "", "pivotal tracker api key")
	listenHost         = envflag.String("LISTEN_HOST", DEFAULT_LISTEN_HOST, fmt.Sprintf("interface to listen on, default: %v", DEFAULT_LISTEN_HOST))
	listenPort         = envflag.Int("LISTEN_PORT", DEFAULT_LISTEN_PORT, fmt.Sprintf("port to listen on, default: %v", DEFAULT_LISTEN_PORT))
	tlsCertificatePath = envflag.String("TLS_CERTIFICATE_PATH", DEFAULT_TLS_CERTIFICATE_PATH, fmt.Sprintf("path to the x509 certificate, default: %v", DEFAULT_TLS_CERTIFICATE_PATH))
	tlsKeyPath         = envflag.String("TLS_KEY_PATH", DEFAULT_TLS_KEY_PATH, fmt.Sprintf("path to x509 private key, default: %v", DEFAULT_TLS_KEY_PATH))
)

func parseFlags() Config {
	envflag.Parse()

	cfg := Config{
		ListenHost:         *listenHost,
		ListenPort:         *listenPort,
		TLSCertificatePath: *tlsCertificatePath,
		TLSKeyPath:         *tlsKeyPath,
		PivotalAPIKey:      *pivotalAPIKey,
	}

	return cfg
}
