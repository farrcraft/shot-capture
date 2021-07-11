package core

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Path     string `json:"-"`
	LogLevel string `json:"log_level"`
	LogPath  string `json:"log_path"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{
		Path: path,
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
