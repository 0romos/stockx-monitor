package config

import (
	"os"
	"strings"
)

var (
	ProductURLs = getEnvAsSlice("PRODUCT_URLS", []string{"https://stockx.com/air-jordan-1-low-paris"})
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}

func getEnvAsSlice(name string, defaultVal []string) []string {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultVal
	}
	return strings.Split(valueStr, ",")
}
