package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Server is a server config structure
type Server struct {
	BindAddr string `yaml:"bindAddr" mapstructure:"bindAddr"`
}

// Keys represents an secret keys configuration
type Keys struct {
	SecretKey  string `yaml:"secret_key" mapstructure:"secret_key"`
	RefreshKey string `yaml:"refresh_key" mapstructure:"refresh_key"`
	TokenTTL   string `yaml:"token_ttl" mapstructure:"token_ttl"`
}

// Settings is a global app settings
type Settings struct {
	DBUrl  string `yaml:"dbUrl" mapstructure:"db_url"`
	Keys   Keys
	Server Server `yaml:"server" mapstructure:"server"`
}

// New returns new application config
func New(path, name, typ string) (
	*Settings, error) {

	cfg := viper.New()

	if path == "" || name == "" || typ == "" {
		return nil, errors.New(
			"no settings file information provided")
	}

	cfg.AddConfigPath(path)
	cfg.SetConfigName(name)
	cfg.SetConfigType(typ)
	if err := cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	s := new(Settings)
	if err := cfg.Unmarshal(s); err != nil {
		return nil, err
	}

	return s, nil
}
