package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	SlackToken     string
	TargetUsers    []string
	TargetChannels []string
	MembersApi     string
	ChannelsApi    string
	InviteApi      string
}

type ConfigFile struct {
	SlackToken     string   `json:"slack_token"`
	TargetUsers    []string `json:"users"`
	TargetChannels []string `json:"channels"`
}

var once sync.Once
var conf *Config

func GetConfig() *Config {
	once.Do(func() {
		var err error
		var cf ConfigFile

		file, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatal("scit: sorry, need config.json!")
		}

		json.Unmarshal(file, &cf)

		conf = &Config{
			SlackToken:     cf.SlackToken,
			TargetUsers:    cf.TargetUsers,
			TargetChannels: cf.TargetChannels,
			MembersApi:     "https://slack.com/api/users.list",
			ChannelsApi:    "https://slack.com/api/channels.list",
			InviteApi:      "https://slack.com/api/channels.invite",
		}
	})

	return conf
}
