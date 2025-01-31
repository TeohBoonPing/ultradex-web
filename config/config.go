package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type FlyConfig struct {
	Env struct {
		DatabaseURL string `toml:"DATABASE_URL"`
	} `toml:"env"`
}

// LoadConfig loads the configuration from the specified toml file.
func LoadConfig() (FlyConfig, error) {
	var flyCfg FlyConfig
	if _, err := toml.DecodeFile("fly.toml", &flyCfg); err != nil {
		return FlyConfig{}, fmt.Errorf("unable to load fly.toml: %v", err)
	}

	return flyCfg, nil
}
