package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type database struct {
	Driver string `json:"driver"`
	Dsn    string `json:"dsn"`
}

type server struct {
	Port               string `json:"port"`
	ReadTimeout        int    `json:"readtimeout"`
	WriteTimeout       int    `json:"writetimeout"`
	MaxHeaderMegabytes int    `json:"maxheadermegabytes"`
}

type token struct {
	Ttl int    `json:"tokenttl"`
	Key string `json:"key"`
}

type Config struct {
	Database database `json:"database"`
	Server   server   `json:"server"`
	Token    token    `json:"token"`
}

func Init(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Println(err, "error")
		return nil, err
	}

	return &config, nil
}
