package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Datebase struct {
	URL  string
	Name string
}

type Config struct {
	DB   Datebase
	Addr string
}

func Load() (*Config, error) {

	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Errorf("unable to open config.json (%v)", err)
		return nil, err
	}

	decoder := json.NewDecoder(file)

	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		fmt.Errorf("Malformated json: %v", err)
		return nil, err
	}
	return config, nil
}
