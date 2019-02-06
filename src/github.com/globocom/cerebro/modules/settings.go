package modules

import (
	"github.com/kelseyhightower/envconfig"
)

// Settings holds configuration data from environment variables.
type Settings struct {
	Debug    bool
	Port     int             `default:"8088"`
	LogLevel LogLevelDecoder `split_words:"true" default:"WARN"`
}

// LoadSettings loads settings from environment variables.
func LoadSettings() (Settings, error) {
	var s Settings
	err := envconfig.Process("", &s)
	return s, err
}
