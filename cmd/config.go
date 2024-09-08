package cmd

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Host     string
	Username string
	Password string
}

func resolveConfigPath() string {
	config_dir := os.Getenv("XDG_CONFIG_HOME")

	if len(config_dir) == 0 {
		config_dir = os.Getenv("HOME") + "/.config/plink"
	}

	return config_dir + "/plink.toml"
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	contents, err := readFile(resolveConfigPath())

	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(contents, config); err != nil {
		return nil, err
	}

	return config, nil
}
