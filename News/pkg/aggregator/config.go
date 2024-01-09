package aggregator

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Feeds   []string `json:"feeds"`
	Refresh int      `json:"refresh"`
}

var configPath string
var feeds []string
var refresh int

func LoadConfig() {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error unmarshalling config:", err)
	}

	feeds = config.Feeds
	refresh = config.Refresh
}
