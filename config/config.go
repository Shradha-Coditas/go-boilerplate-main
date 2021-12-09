package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

//ServerConfig - Server configuration.
type ServerConfig struct {
	Port            int
	AllowOrigins    []string
	LogLevel        string
	DefaultLanguage string
	Languages       []string
}

// MySqlConfig - MySqlConfig configuration
type MySqlConfig struct {
	URL      string
	Database string
}

// RedisConfig - Redis configuration
type RedisConfig struct {
	URL      string
	Password string
	Database int
}

//Config structure.
type Config struct {
	Server ServerConfig
	MySQL  MySqlConfig
	Redis  RedisConfig
}

//AppConfig - Appconfig object,.
var AppConfig = &Config{
	Server: ServerConfig{
		Port:            3000,
		AllowOrigins:    []string{"*"},
		LogLevel:        "info",
		DefaultLanguage: "en",
		Languages:       []string{"en"},
	},
	MySQL: MySqlConfig{
		URL:      "",
		Database: "",
	},
	Redis: RedisConfig{
		URL:      "",
		Password: "",
		Database: 0,
	},
}

// LoadEnv - function load Enviroment variable from .env file.
func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		Server: ServerConfig{
			Port:            getEnvAsInt("API_PORT", 3000),
			AllowOrigins:    strings.Split(getEnv("ALLOW_ORIGIN", "*"), ","),
			LogLevel:        getEnv("LOG_LEVEL", "info"),
			DefaultLanguage: getEnv("DEFAULT_LANGUAGE", "en"),
			Languages:       strings.Split(getEnv("LANGUAGES", "en"), ","),
		},
		MySQL: MySqlConfig{
			URL:      getEnv("URL", ""),
			Database: getEnv("DATABASE", ""),
		},
		Redis: RedisConfig{
			URL:      getEnv("REDIS_URL", ""),
			Password: getEnv("REDIS_PASSWORD", ""),
			Database: getEnvAsInt("REDIS_DATABASE", 0),
		},
	}
	AppConfig = config
	return config
}

// Simple helper function to read an environment or return a default value.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
