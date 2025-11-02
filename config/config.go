package config

import (
	"log"
	"strings"
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
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetDefault("server.port", "8080")
		viper.SetDefault("env", "dev")

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Config could not be loaded, please validate that config file is present")
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
