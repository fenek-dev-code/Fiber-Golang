package config

import (
	"os"
	"strconv"

	"log"
)

func getIntEnv(key string, defaultValue int) int {
	envValue := os.Getenv(key)
	value, err := strconv.Atoi(envValue)
	if err != nil {
		log.Println("Failed to parse int from env, using default:", key, err)
		return defaultValue
	}
	return value
}
func getStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Println("Failed to parse string from env, using default:", key, value)
		return defaultValue
	}
	return value
}

func getBoolEnv(key string, defaultValue bool) bool {
	envValue := os.Getenv(key)
	value, err := strconv.ParseBool(envValue)
	if err != nil {
		log.Println("Failed to parse bool from env, using default:", key, err)
		return defaultValue
	}
	return value
}
