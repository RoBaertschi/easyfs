package config

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

type Config struct {
	ServeDirectory fs.FS
}

type Filesystem struct {
	ServeDirectory string `mapstructure:"serve_directory"`
}

func ReadConfig(path string) error {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(toml.Driver)

	err := config.LoadFiles(path)
	if err != nil {
		return err
	}

	filesystem := Filesystem{}
	err = config.BindStruct("fs", &filesystem)

	if err != nil {
		return err
	}

	absoluteServeDir, err := filepath.Abs(filesystem.ServeDirectory)
	if err != nil {
		return err
	}

}
