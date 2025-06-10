package config

import (
	"os"
	"sync"
)

type Configuration struct {
	Port                  string
	Spotify_client_ID     string
	Spotify_client_secret string
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func buildConfig() *Configuration {
	return &Configuration{
		Port:                  getEnv("PORT", "3000"),
		Spotify_client_ID:     getEnv("SPOTIFY_CLIENT_ID", ""),
		Spotify_client_secret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
	}
}

var configuration *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		configuration = buildConfig()
	})
	return configuration
}
