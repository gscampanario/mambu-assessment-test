package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Env string `mapstructure:"env"`
}

var (
	cfg  *Config
	once sync.Once
)

// Load loads information from config.yaml into a singleton, used for environment related parameters.
func Load() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})
}

// Get returns the singleton Config instance.
func Get() *Config {
	if cfg == nil {
		Load()
	}
	return cfg
}
