package application

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort: 3000,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	if serverPortStr, exists := os.LookupEnv("SERVER_PORT"); exists {
		if serverPort, err := strconv.ParseUint(serverPortStr, 10, 16); err == nil {
			cfg.ServerPort = uint16(serverPort)
		}
	}

	return cfg
}