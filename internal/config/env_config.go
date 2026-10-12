package config

import (
	"fmt"
	"os"
	"strconv"
)

func LoadSingleEnvVar[T string | int | float32 | bool](key string, defaultValue T) T {
	val, isExists := os.LookupEnv(key)

	if !isExists {
		return defaultValue
	}

	switch any(defaultValue).(type) {
	case string:
		return any(val).(T)
	case int:
		f, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultValue
		}

		return any(f).(T)
	case float32:
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultValue
		}

		return any(float32(f)).(T)

	case bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultValue
		}

		return any(b).(T)

	default:
		fmt.Println("Unsupported datatype has been passed for the env var: " + key)
		return defaultValue
	}
}
