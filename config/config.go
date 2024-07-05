package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

type Config struct {
	ServeDirectory string
}

type Filesystem struct {
	ServeDirectory string `mapstructure:"serve_directory" default:"./static"`
}

func ReadConfig(path string) (*Config, error) {
	config.WithOptions(config.ParseEnv, config.ParseDefault)
	config.AddDriver(toml.Driver)

	err := config.LoadFiles(path)
	if err != nil {
		return nil, err
	}

	filesystem := Filesystem{}
	err = config.BindStruct("fs", &filesystem)

	if err != nil {
		return nil, fmt.Errorf("while reading config tag \"fs\": %s", err.Error())
	}

	absoluteServeDir, err := filepath.Abs(filesystem.ServeDirectory)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(absoluteServeDir); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("The provided serve dir %s does not exist!", absoluteServeDir)
	}
	return &Config{
		ServeDirectory: absoluteServeDir,
	}, nil
}
