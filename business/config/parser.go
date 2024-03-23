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

func StoreConfig(path string, cfg Config) error {
	bs, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, bs, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
