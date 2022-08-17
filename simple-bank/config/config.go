package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost        string        `mapstructure:"POSTGRES_HOST"`
	PostgresUser        string        `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string        `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB          string        `mapstructure:"POSTGRES_DB"`
	PostgresPort        string        `mapstructure:"POSTGRES_PORT"`
	PostgresSSLMode     string        `mapstructure:"POSTGRES_SSL_MODE"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// Note, these are necessary for the struct marshalling to pick up env variables in the case that the config file does not exist
func setEnvDefaults() {
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "")
	viper.SetDefault("POSTGRES_DB", "simple-bank")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_SSL_MODE", "disable")
	viper.SetDefault("TOKEN_SYMMETRIC_KEY", "")
	viper.SetDefault("ACCESS_TOKEN_DURATION", "10m")
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	// This is how you would specify the name of the config file, ex app.env
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// If the env variables exist in the env, they will overwrite the file values
	viper.AutomaticEnv()
	setEnvDefaults()

	err := viper.BindEnv("POSTGRES_DB")

	if err != nil {
		log.Println("error binding env var")
	}
	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Attempt to read from env variables
			log.Println(".env file not found, hopefully you set the env variables!")
		} else {
			return nil, err
		}
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	log.Println("config db")
	log.Println(config.PostgresDB)
	log.Println("env db")
	log.Println(os.Getenv("POSTGRES_DB"))
	return &config, nil
}
