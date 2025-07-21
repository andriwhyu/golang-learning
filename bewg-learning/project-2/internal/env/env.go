package env

import (
	"os"
	"strconv"
	"time"
)

func GetStringVar(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetIntVar(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetDurationStringVar(key string, fallback string) time.Duration {
	val, ok := os.LookupEnv(key)

	if !ok {
		duration, err := time.ParseDuration(fallback)
		if err != nil {
			return 15 * time.Minute
		}
		return duration
	}

	duration, err := time.ParseDuration(val)
	if err != nil {
		return 15 * time.Minute
	}

	return duration
}
