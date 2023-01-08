package env

import (
	"github.com/joho/godotenv"
	zeroLog "github.com/rs/zerolog/log"
	"trading-bot/internal/config"
)

// TODO
func Read(envFile ...string) config.Config {
	_, err := godotenv.Read(envFile...)
	if err != nil {
		zeroLog.Fatal().Err(err)
	}

	//r := reader(envs)

	return config.Config{}
}
