package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"trading-bot/internal/config"
)

// TODO
func Read(envFile ...string) config.Config {
	_, err := godotenv.Read(envFile...)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	//r := reader(envs)

	return config.Config{}
}
