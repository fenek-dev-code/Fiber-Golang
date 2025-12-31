package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(confPath ...string) {
	if len(confPath) == 0 {
		confPath = append(confPath, ".env")
	}
	for _, path := range confPath {
		if err := godotenv.Load(path); err != nil {
			log.Printf("Error loading %s file", path)
			continue
		}
		log.Printf("Loaded %s config file", path)
	}
}

type LogConfig struct {
	Level  int
	Format string
}

type EnvConfig struct {
	DatabaseURL string
	*LogConfig
}

func GetEnvConfig() *EnvConfig {
	return &EnvConfig{
		DatabaseURL: getStringEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname"),
		LogConfig: &LogConfig{
			Level:  getIntEnv("LOG_LEVEL", 3),
			Format: getStringEnv("LOG_FORMAT", "json"),
		},
	}
}
