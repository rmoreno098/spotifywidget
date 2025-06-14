package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Configuration struct {
	Port                string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyRedirectUri  string
	JWTSecret           string
	RedisAddr           string
	RedisPassword       string
	RedisDb             string
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if fallback == "" {
		panic("Environment variable " + key + " not defined")
	}

	return fallback
}

func buildConfig() *Configuration {
	_ = godotenv.Load()
	return &Configuration{
		Port:                getEnv("PORT", "3000"),
		SpotifyClientId:     getEnv("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyRedirectUri:  getEnv("SPOTIFY_REDIRECT_URI", ""),
		JWTSecret:           getEnv("JWT_SECRET", ""),
		RedisAddr:           getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:       getEnv("REDIS_PASSWORD", ""),
		//RedisDb:             getEnv("REDIS_DB", "0"),
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
