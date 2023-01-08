package main

import (
	zeroLog "github.com/rs/zerolog/log"
	"trading-bot/internal/logger/console"
	"trading-bot/internal/platform/placeholder"
	"trading-bot/internal/strategy/basic"
)

func main() {
	platform := placeholder.New()
	logger := console.New()

	strategy := basic.New(platform, logger)

	err := <-strategy.Run()
	if err != nil {
		zeroLog.Error().Err(err)
	}
}