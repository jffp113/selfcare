package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ConfigFromFile(path string) (Config, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading db file: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return cfg, nil
	}

	return cfg, nil
}
