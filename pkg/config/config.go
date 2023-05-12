package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CurrentBoard string `yaml:"current_board"`
}

func New(b []byte) (Config, error) {
	var conf Config
	err := yaml.Unmarshal(b, &conf)
	return conf, err
}

func (c Config) Write(w io.Writer) error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
