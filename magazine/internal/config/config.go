package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application Application `yaml:"application" env-required:"true"`
	DB          DB          `yaml:"db" env-required:"true"`
}

type Application struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

type DB struct {
	Host     string `yaml:"host" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	DBname   string `yaml:"dbname" env-required:"true"`
}

var (
	config *Config
	once   sync.Once
)

func MustLoad() *Config {
	if config == nil {
		once.Do(
			func() {
				configPath := filepath.Join("..", "..", "config.yaml")

				if _, err := os.Stat(configPath); err != nil {
					log.Fatalf("Error opening config file: %s", err)
				}

				var newConfig Config
				err := cleanenv.ReadConfig(configPath, &newConfig)
				if err != nil {
					log.Fatalf("Error reading config file: %s", err)
				}

				config = &newConfig
			})
	}

	return config
}
