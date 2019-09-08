package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type Config struct {
	path string
	Data ConfigData
}

type ConfigData struct {
	Userconfig *user.Users         `json:"users"`
	WAClients  *waclient.WAClients `json:"clients"`
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	config.path = path
	fi, err := os.Stat(config.path)
	if err != nil {
		config.init()
		if err := config.Save(); err != nil {
			return nil, err
		}
		return config, nil
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("NotADirectoryError: %s is not a valid file", config.path)
	}
	if err := config.load(); err != nil {
		config.init()
		return nil, err
	}
	config.init()
	return config, nil
}

func (c *Config) load() error {
	file, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c.Data)
	for k := range c.Data.WAClients.Clients {
		c.Data.WAClients.Clients[k].ImportSession()
	}
	return nil
}

func (c *Config) Save() error {
	for k := range c.Data.WAClients.Clients {
		c.Data.WAClients.Clients[k].ExportSession()
	}
	datajson, err := json.MarshalIndent(c.Data, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.path, datajson, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) init() {
	if c.Data.WAClients == nil {
		c.Data.WAClients = new(waclient.WAClients)
	}
	if c.Data.WAClients.Clients == nil {
		c.Data.WAClients.Clients = make(map[string]*waclient.WAClientConfig)
	}
	if c.Data.Userconfig == nil {
		c.Data.Userconfig = user.NewUsers()
	}
	for n := range *c.Data.Userconfig.Users {
		if (*c.Data.Userconfig.Users)[n].Clients == nil {
			(*c.Data.Userconfig.Users)[n].Clients = &[]string{}
		}
	}
}
