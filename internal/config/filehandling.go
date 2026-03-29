package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	filename = ".gatorconfig.json"
)

func Read() (Config, error) {
	config := Config{}

	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory: %v", err)
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("config file " + path + " not found.")
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			return Config{}, fmt.Errorf("error reading file; %v", err)
		}
		err = json.Unmarshal(data, &config)
		if err != nil {
			return Config{}, fmt.Errorf("error unmarshaling json; %v", err)
		}
	}
	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.Current_user_name = user
	err := writeConfigFile(*cfg)
	if err != nil {
		return fmt.Errorf("error writing user to file; %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + string(os.PathSeparator) + filename, nil

}

func writeConfigFile(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
