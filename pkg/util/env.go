package util

import (
	"os"
	"strconv"
)

func EnvStr(key string, defaultValue string) string {
	if v, exist := os.LookupEnv(key); exist {
		return v
	}

	return defaultValue
}

func EnvInt(key string, defaultValue int) int {
	if v, exist := os.LookupEnv(key); exist {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}

	return defaultValue
}

func EnvBool(key string, defaultValue bool) bool {
	if v, exist := os.LookupEnv(key); exist {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}

	return defaultValue
}
