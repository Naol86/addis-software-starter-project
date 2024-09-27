package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBUri			string `mapstructure:"DB_URI"`
	DBName     string `mapstructure:"DB_NAME"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiryHour int `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	ContextTimeout int `mapstructure:"CONTEXT_TIMEOUT"`
}

// NewEnv creates and loads environment variables from .env file
func NewEnv() (*Env, error) {
	env := Env{}

	// Configure viper to read .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config into the Env struct
	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %w", err)
	}

	// Check if the application is in development mode
	if env.AppEnv == "development" {
		log.Println("App is running in development mode")
	}

	return &env, nil
}