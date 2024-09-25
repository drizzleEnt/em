package env

import (
	"errors"
	"net"
	"os"
)

const (
	httpHostName = "HTTP_HOST"
	httpPortName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}
	port := os.Getenv(httpPortName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
