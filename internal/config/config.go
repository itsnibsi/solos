package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Bind string `mapstructure:"BIND"`
	Port int    `mapstructure:"PORT"`
	Env  string `mapstructure:"ENV"`
}

func Load() (*Config, error) {
	viper.SetDefault("BIND", "0.0.0.0")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("ENV", "development")

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
