package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	envConfigPath = "CONFIG_PATH"
	envConfigName = "CONFIG_NAME"
)

type Config struct {
	Database Postgres `yaml:"database"`
}

// Interface is used for easy passing of types into different internal packages
type Interface interface {
	Database
}

// New creates a new config Interface
func New() (Interface, error) {
	configPath, ok := os.LookupEnv(envConfigPath)
	if !ok {
		return nil, fmt.Errorf("missing %s env", envConfigPath)
	}

	configName, ok := os.LookupEnv(envConfigName)
	if !ok {
		return nil, fmt.Errorf("missing %s env", envConfigName)
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	// defaults for optional config
	viper.SetDefault("database.poolMaxConns", 10)
	viper.SetDefault("database.connectTimeout", 10)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	cfg := Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshaling: %w", err)
	}

	return cfg, nil
}
