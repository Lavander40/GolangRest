package apiserver

import "github.com/lavander40/golang_rest/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":6070",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
