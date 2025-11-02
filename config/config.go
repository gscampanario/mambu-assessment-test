package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
		Auth struct {
			Username string
			Password string
		}
	}
	Env     string
	Storage struct {
		BankFolder        string
		DBFileLocation    string
		TxnFeedbackFolder string
	}
}

var (
	cfg  *Config
	once sync.Once
)

// Load loads information from .env into a singleton, used for environment related parameters.
func Load() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		cfg = &Config{
			Server: struct {
				Port string
				Auth struct {
					Username string
					Password string
				}
			}{
				Port: os.Getenv("SERVER_PORT"),
				Auth: struct {
					Username string
					Password string
				}{
					Username: os.Getenv("BASIC_AUTH_USERNAME"),
					Password: os.Getenv("BASIC_AUTH_PASSWORD"),
				},
			},
			Env: os.Getenv("ENVIRONMENT"),
			Storage: struct {
				BankFolder        string
				DBFileLocation    string
				TxnFeedbackFolder string
			}{
				BankFolder:        os.Getenv("BANK_FOLDER"),
				DBFileLocation:    os.Getenv("SQLITE_DB_FILE_LOCATION"),
				TxnFeedbackFolder: os.Getenv("TXN_FBK_FOLDER"),
			},
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
