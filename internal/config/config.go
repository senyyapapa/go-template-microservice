package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Logger struct {
		Level        *slog.Level `yaml:"level"`
		ShowPathCall bool        `yaml:"show_path_call" env-default:"false"`
	} `yaml:"logger"`
	// Other config fields
}

func MustLoad() *Config {
	godotenv.Load()
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(configPath); err != nil {
		panic(err)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic(err.Error())
	}
	return cfg

}

func fetchConfigPath() (res string) {
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "config/config_local.yaml"
	}
	return
}
