package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	Driver string `json:"driver"`
	Dsn    string `json:"dsn"`
}

type Server struct {
	Port               string `json:"port"`
	ReadTimeout        int    `json:"readtimeout"`
	WriteTimeout       int    `json:"writetimeout"`
	MaxHeaderMegabytes int    `json:"maxheadermegabytes"`
}

type Token struct {
	Ttl int    `json:"tokenttl"`
	Key string `json:"key"`
}

type Hash struct {
	Cost int `json:"cost"`
}

type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
	Token    Token    `json:"token"`
	Hash     Hash     `json:"hash"`
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
