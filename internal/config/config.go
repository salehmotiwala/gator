package config

import (
	"encoding/json"
	"io/fs"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	path, err := getConfigPath()

	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	path, err := getConfigPath()

	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func getConfigPath() (string, error) {
	path, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path + ("/" + configFileName), nil
}
