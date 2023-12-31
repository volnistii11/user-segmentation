package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	DatabaseDSN       string `yaml:"database_dsn"`
	DatabaseDriver    string `yaml:"database_driver"`
	HTTPServerAddress string `yaml:"http_server_address"`
}

func New() *Config {
	return &Config{
		DatabaseDSN:       "",
		DatabaseDriver:    "",
		HTTPServerAddress: "",
	}
}

func (cfg *Config) Parse() error {
	defaultValues, err := parseDefaultValues()
	if err != nil {
		return errors.Wrap(err, "Parse.parseDefaultValues")
	}

	flag.StringVar(&cfg.DatabaseDSN, "d", defaultValues.DatabaseDSN, "database dsn")
	flag.StringVar(&cfg.DatabaseDriver, "r", defaultValues.DatabaseDriver, "database driver")
	flag.StringVar(&cfg.HTTPServerAddress, "a", defaultValues.HTTPServerAddress, "http router address")
	flag.Parse()

	return nil
}

func parseDefaultValues() (*Config, error) {
	configPath := "../../config/config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "osStat")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "ReadConfig")
	}

	return &cfg, nil
}
