package config

import "github.com/ilyakaznacheev/cleanenv"

type ServerConfig struct {
	IsServer bool `yaml:"is_server" env:"IS_SERVER" env-default:"false"`
}

type ClientConfig struct {
	WebServer string `yaml:"web_server" env:"WEB_SERVER" env-default:"net"`
	Port      string `yaml:"port" env:"PORT" env-default:"3000"`
}

type Config struct {
	Server         ServerConfig `yaml:"server"`
	EmitterAddress string       `yaml:"emitter_address" env:"EMITTER_ADDRESS" env-default:"localhost:1234"`
	Client         ClientConfig `yaml:"client"`
}

var cfg Config

func GetConfig() *Config {
	return &cfg
}

func ReadConfig(path string) (*Config, error) {
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
