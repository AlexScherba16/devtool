package config

import "time"

type ServerConfig struct {
	Address      string
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
