package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	path string
	data ConfigData
}

type ConfigData struct {
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("NotADirectoryError: %s is not a valid file", path)
	}
	config.path = path
	if err := config.load(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) load() error {
	file, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c.data)
	return nil
}

func (c *Config) Save() error {
	datajson, err := json.Marshal(c.data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.path, datajson, 0644)
	if err != nil {
		return err
	}
	return nil
}
