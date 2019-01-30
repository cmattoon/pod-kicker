package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const (
	ENV_KEY_TOKEN = "PK_TOKEN"
	ENV_KEY_PORT  = "PK_PORT"
)

type Config struct {
	Token      string
	ListenPort int
}

func (c *Config) GetListenPort() string {
	return fmt.Sprintf(":%d", c.ListenPort)
}

func NewConfig(argv []string) *Config {
	cfg := &Config{
		Token:      "super-secret",
		ListenPort: 4200,
	}

	if uid, err := uuid.NewUUID(); err == nil {
		cfg.Token = uid.String()
	}

	// Move to kingpin when expanding
	// Set token
	if token, ok := os.LookupEnv(ENV_KEY_TOKEN); ok {
		if token != "" {
			cfg.Token = token
		}
	}

	if port, ok := os.LookupEnv(ENV_KEY_PORT); ok {
		p, err := strconv.Atoi(port)
		if err != nil {
			cfg.ListenPort = p
		}
	}

	return cfg
}
